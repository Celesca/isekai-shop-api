package service

import (
	_itemManagingRepository "github.com/Celesca/isekai-shop-api/pkg/itemManaging/repository"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository}
}
