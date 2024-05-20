package router

import (
	"errors"
	"evelp/model"
	"evelp/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func order(c *gin.Context) {
	regionId, err := strconv.Atoi(c.Query("regionId"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	scope, err := strconv.ParseFloat(c.Query("scope"), 64)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	isBuyOrder, err := strconv.ParseBool(c.Query("isBuyOrder"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	lang := c.Query("lang")
	if lang == "" {
		c.AbortWithError(500, errors.New("lang is empty"))
		return
	}

	offerId, err := strconv.Atoi(c.Query("offerId"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	offer, err := model.GetOffer(offerId)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	itemId := offer.ItemId
	isBluePrint := offer.IsBluePrint

	orderService := service.NewOrderService(itemId, regionId, isBluePrint, float64(scope))
	orders, err := orderService.Orders(isBuyOrder, lang)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, orders)
}
