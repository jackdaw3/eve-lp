package router

import (
	"errors"
	"evelp/model"
	"evelp/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func region(c *gin.Context) {
	factions, err := model.GetRegions()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, factions)
}

func regionDetail(c *gin.Context) {
	regionId, err := strconv.Atoi(c.Query("regionId"))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	lang := c.Query("lang")
	if lang == "" {
		c.AbortWithError(500, errors.New("lang is empty"))
		return
	}

	regionService := service.NewRegionService(lang)
	regionDTO, err := regionService.Region(regionId)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, regionDTO)
}
