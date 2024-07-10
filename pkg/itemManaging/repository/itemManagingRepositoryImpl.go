package repository

import (
	"fmt"

	"github.com/Celesca/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemManagingException "github.com/Celesca/isekai-shop-api/pkg/itemManaging/exception"
	_itemManagingModel "github.com/Celesca/isekai-shop-api/pkg/itemManaging/model"
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

	if err := r.db.Create(itemEntity).Scan(item).Error; err != nil {
		r.logger.Errorf("Creating item failed: %s", err.Error())
		return nil, &_itemManagingException.ItemCreating{}
	}

	fmt.Print(item)

	return item, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	if err := r.db.Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error; err != nil {
		r.logger.Errorf("Editing item failed: %s", err.Error())
		return 0, &_itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID uint64) error {
	if err := r.db.Table("items").Where("id = ?", itemID).Update("is_archive", true).Error; err != nil {
		r.logger.Errorf("Archiving item failed: %s", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}