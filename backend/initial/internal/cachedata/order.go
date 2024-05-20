package cachedata

import (
	"encoding/json"
	"evelp/config/global"
	"evelp/log"
	"evelp/model"
	"evelp/util/cache"
	"evelp/util/net"
	"fmt"

	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

type orderData struct {
	orders         map[string]*model.Orders
	expirationTime time.Duration
	items          map[int]interface{}
}

func (o *orderData) Refresh() error {
	regions := []int{THE_FORGE, DOMAIN, SINQ_LAISION, HEIMATER}
	go func() {
		for {
			for _, region := range regions {
				if err := o.refreshOrdersByRegion(region); err != nil {
					log.Errorf(err, "failed to load orders from region %d", region)
				}
			}
		}
	}()

	return nil
}

func (o *orderData) refreshOrdersByRegion(regionId int) error {
	log.Debugf("loading orders from region %d", regionId)
	if err := o.ordersByRegion(regionId); err != nil {
		o.clearMap()
		return errors.WithMessagef(err, "failed to load orders from region %d", regionId)
	}

	log.Debugf("saving orders from region %d to redis", regionId)
	for key, order := range o.orders {
		if err := cache.Set(key, *order, o.expirationTime); err != nil {
			log.Errorf(err, "failed to save order %s to redis", key)
		}
	}

	log.Debugf("saved orders from region %d to redis", regionId)
	o.clearMap()

	return nil
}

func (o *orderData) ordersByRegion(regionId int) error {
	pages, err := o.marketPages(regionId)
	if err != nil {
		return err
	}

	for i := 1; i <= pages; i++ {
		wg.Add(1)
		global.Ants.Submit(o.ordersByRegionPage(regionId, i))
	}

	wg.Wait()
	return nil
}

func (o *orderData) ordersByRegionPage(regionId int, page int) func() {
	return func() {
		defer wg.Done()

		req := fmt.Sprintf("%s/markets/%d/orders/?datasource=%s&order_type=all&page=%d",
			global.Conf.Data.Remote.Address,
			regionId,
			global.Conf.Data.Remote.DataSource,
			page,
		)

		content, _, err := net.Get(req)
		if err != nil {
			log.Errorf(err, "failed to get %d region %d page's orders", regionId, page)
			return
		}

		var orders model.Orders
		if err = json.Unmarshal(content, &orders); err != nil {
			log.Errorf(err, "failed to unmarshal %d region %d page's orders json", regionId, page)
			return
		}

		for _, order := range orders {
			_, ok := o.items[order.ItemId]
			if !ok {
				continue
			}
			key := cache.Key("order", strconv.Itoa(regionId), strconv.Itoa(order.ItemId))
			o.syncPutToMap(key, &order)
		}
	}
}

func (o *orderData) marketPages(regionId int) (int, error) {
	req := fmt.Sprintf("%s/markets/%d/orders/?datasource=%s&order_type=all&page=%d",
		global.Conf.Data.Remote.Address,
		regionId,
		global.Conf.Data.Remote.DataSource,
		1,
	)

	_, header, err := net.Get(req)
	if err != nil {
		return 0, err
	}

	pages, err := strconv.Atoi(header.Get("x-pages"))
	if err != nil {
		return 0, err
	}

	return pages, nil

}

func (o *orderData) syncPutToMap(key string, order *model.Order) {
	defer mu.Unlock()
	mu.Lock()

	orders, ok := o.orders[key]
	order.LastUpdated = time.Now()
	if ok {
		val := append(*orders, *order)
		o.orders[key] = &val
	} else {
		val := model.Orders{*order}
		o.orders[key] = &val
	}
}

func (o *orderData) clearMap() {
	for k := range o.orders {
		delete(o.orders, k)
	}
}
