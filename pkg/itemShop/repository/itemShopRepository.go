package repository

import (
	"github.com/Celesca/isekai-shop-api/entities"
	_itemShopModel "github.com/Celesca/isekai-shop-api/pkg/itemShop/model"
)

// ที่เป็น Interface เพราะ ถ้าสมมติเราจะสร้างเป็น Mock
type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
}
