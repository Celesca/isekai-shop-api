package repository

import (
	"github.com/Celesca/isekai-shop-api/databases"
	"github.com/Celesca/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(
	db databases.Database,
	logger echo.Logger,
) AdminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *adminRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	return nil, nil
}

func (r *adminRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	return nil, nil
}
