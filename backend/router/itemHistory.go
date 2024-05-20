package router

import (
	"evelp/model"
	"evelp/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func itemHistory(c *gin.Context) {
	regionId, err := strconv.Atoi(c.Query("regionId"))
	if err != nil {
		c.AbortWithError(500, err)
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

	itemHistoryService := service.NewItemHistoryService(itemId, regionId, isBluePrint)
	itemHistorys, err := itemHistoryService.History()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, itemHistorys)
}
