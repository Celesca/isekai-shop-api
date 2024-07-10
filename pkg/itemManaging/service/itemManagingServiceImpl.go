package service

import (
	"github.com/Celesca/isekai-shop-api/entities"
	_itemManagingModel "github.com/Celesca/isekai-shop-api/pkg/itemManaging/model"
	_itemManagingRepository "github.com/Celesca/isekai-shop-api/pkg/itemManaging/repository"
	_itemShopModel "github.com/Celesca/isekai-shop-api/pkg/itemShop/model"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	itemShopRepository _itemManagingRepository.ItemShopRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{
		itemManagingRepository,
		itemShopRepository,
	}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	// Entity to send to repository
	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	itemEntityResult, err := s.itemManagingRepository.Editing(itemID, itemEditingReq)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}
