package service

import (
	_itemShopModel "github.com/Celesca/isekai-shop-api/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*_itemShopModel.Item, error)
}
