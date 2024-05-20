package service

import (
	"evelp/model"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestRegion(t *testing.T) {
	defer monkey.UnpatchAll()
	region := mockRegionData()
	mockRegion(region)

	regionService := NewRegionService("en")
	regionDTO, err := regionService.Region(10000001)

	assert.NoError(t, err)
	assert.Equal(t, region.Name.En, regionDTO.RegionName)
}

func mockRegionData() *model.Region {
	return &model.Region{
		RegionId: 10000001,
		Name:     model.Name{En: "Derelik"},
	}
}

func mockRegion(region *model.Region) {
	monkey.Patch(model.GetRegion, func(regionId int) (*model.Region, error) {
		return region, nil
	})
}
