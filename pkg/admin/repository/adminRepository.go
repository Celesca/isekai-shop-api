package repository

import "github.com/Celesca/isekai-shop-api/entities"

type AdminRepository interface {
	Creating(adminEntity *entities.Player) (*entities.Player, error)
	FindByID(adminID string) (*entities.Player, error)
}
