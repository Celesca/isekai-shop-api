package repository

import "github.com/Celesca/isekai-shop-api/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindByID(playerID string) (*entities.Player, error)
}
