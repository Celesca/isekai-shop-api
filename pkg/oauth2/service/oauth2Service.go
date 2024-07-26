package service

import (
	_adminModel "github.com/Celesca/isekai-shop-api/pkg/admin/model"
	_playerModel "github.com/Celesca/isekai-shop-api/pkg/player/model"
)

type OAuth2Service interface {
	PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error
}
