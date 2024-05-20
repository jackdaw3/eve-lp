package service

import (
	"evelp/model"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestItem(t *testing.T) {
	defer monkey.UnpatchAll()
	item := mockItemData()
	mockItem(item)

	itemSerivce := NewItemService("en")
	result, err := itemSerivce.Item(34)

	assert.NoError(t, err)
	assert.Equal(t, item.Name.En, result.ItemName)
}

func TestItemDetail(t *testing.T) {
	defer monkey.UnpatchAll()
	item := mockItemData()
	mockItem(item)

	itemSerivce := NewItemService("en")
	result, err := itemSerivce.ItemDetail(34)

	assert.NoError(t, err)
	assert.Equal(t, item.Name.En, result.ItemName)
	assert.Equal(t, item.Volume, result.Volume)
}

func ItemByOffer(t *testing.T) {
	defer monkey.UnpatchAll()
	item := mockItemData()
	mockItem(item)

	itemSerivce := NewItemService("en")
	result, err := itemSerivce.ItemByOffer(340)

	assert.NoError(t, err)
	assert.Equal(t, item.Name.En, result.ItemName)
}

func mockItemData() *model.Item {
	return &model.Item{
		ItemId: 34,
		Name:   model.Name{En: "Tritanium"},
		Volume: 0.01,
	}
}

func mockItem(item *model.Item) {
	monkey.Patch(model.GetItem, func(id int) (*model.Item, error) {
		return item, nil
	})

	monkey.Patch(model.GetItemByOffer, func(offerId int) (*model.Item, error) {
		return item, nil
	})
}
