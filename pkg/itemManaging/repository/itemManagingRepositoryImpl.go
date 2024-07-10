package repository

import (
	"github.com/Celesca/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemManagingException "github.com/Celesca/isekai-shop-api/pkg/itemManaging/exception"
)

type itemManagingRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db *gorm.DB, logger echo.Logger) *itemManagingRepositoryImpl {
	return &itemManagingRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {

	// allocate new item to return
	item := new(entities.Item)

	if err := r.db.Create(itemEntity).Error; err != nil {
		r.logger.Errorf("Creating item failed: %v", err)
		return nil, &_itemManagingException.ItemCreating{}
	}

	return item, nil
}
