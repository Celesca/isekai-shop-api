package repository

import "github.com/Celesca/isekai-shop-api/entities"

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
}
