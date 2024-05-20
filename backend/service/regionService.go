package service

import (
	"evelp/dto"
	"evelp/model"
)

type RegionService struct {
	lang string
}

func NewRegionService(lang string) *RegionService {
	return &RegionService{lang}
}

func (is *RegionService) Region(regionId int) (*dto.RegionDTO, error) {
	var regionDTO dto.RegionDTO

	region, err := model.GetRegion(regionId)
	if err != nil {
		return nil, err
	}

	regionDTO.RegionId = region.RegionId
	regionDTO.RegionName = region.Name.Lang(is.lang)

	return &regionDTO, nil
}
