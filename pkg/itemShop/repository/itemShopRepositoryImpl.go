package repository

import (
	"github.com/Celesca/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopException "github.com/Celesca/isekai-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/Celesca/isekai-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{db, logger}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{}) // select * from items

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}

	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	if err := query.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}
