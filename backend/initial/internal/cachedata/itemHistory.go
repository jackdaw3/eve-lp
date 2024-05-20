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
	"time"

	"github.com/robfig/cron/v3"
)

type itemHistroy struct {
	expirationTime time.Duration
	products       map[int]interface{}
}

func (i *itemHistroy) Refresh() error {
	cron := cron.New(cron.WithSeconds())
	if _, err := cron.AddFunc("@daily", i.invoke()); err != nil {
		return err
	}
	cron.Start()
	return nil
}

func (i *itemHistroy) invoke() func() {
	return func() {
		regions := []int{THE_FORGE, DOMAIN, SINQ_LAISION, HEIMATER}
		for _, region := range regions {
			i.refreshHistoryByRegion(region)
			time.Sleep(time.Minute * 25)
		}
	}
}

func (i *itemHistroy) refreshHistoryByRegion(regionId int) {
	log.Debugf("loading history from region %d", regionId)

	for p := range i.products {
		req := fmt.Sprintf("%s/markets/%d/history/?datasource=%s&type_id=%d",
			global.Conf.Data.Remote.Address,
			regionId,
			global.Conf.Data.Remote.DataSource,
			p,
		)

		content, _, err := net.Get(req)
		if err != nil {
			log.Debugf("failed to get item %d histroy from from region %d", p, regionId)
			continue
		}

		var itemHistorys model.ItemHistorys
		if err = json.Unmarshal(content, &itemHistorys); err != nil {
			log.Errorf(err, "failed to unmarshal item %d history %s json from region %d", p, content, regionId)
			continue
		}

		for _, itemitemHistory := range itemHistorys {
			itemitemHistory.ItemId = p
		}

		key := cache.Key("history", strconv.Itoa(regionId), strconv.Itoa(p))
		if err := cache.Set(key, itemHistorys, i.expirationTime); err != nil {
			log.Errorf(err, "failed to save history %v to redis", key)
		}
	}

	log.Debugf("saved history from region %d to redis", regionId)
}
