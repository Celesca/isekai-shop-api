package service

import (
	"github.com/Celesca/isekai-shop-api/entities"
	_adminRepository "github.com/Celesca/isekai-shop-api/pkg/admin/repository"
	_playerRepository "github.com/Celesca/isekai-shop-api/pkg/player/repository"

	_adminModel "github.com/Celesca/isekai-shop-api/pkg/admin/model"
	_playerModel "github.com/Celesca/isekai-shop-api/pkg/player/model"
)

type googleOAuth2Service struct {
	playerRepository _playerRepository.PlayerRepository
	adminRepository  _adminRepository.AdminRepository
}

func NewGoogleOAuth2Service(
	playerRepository _playerRepository.PlayerRepository,
	adminRepository _adminRepository.AdminRepository,
) OAuth2Service {
	return &googleOAuth2Service{
		playerRepository,
		adminRepository,
	}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *_playerModel.PlayerCreatingReq) error {
	playerEntity := &entities.Player{
		ID:     playerCreatingReq.ID,
		Name:   playerCreatingReq.Name,
		Email:  playerCreatingReq.Email,
		Avatar: playerCreatingReq.Avatar,
	}

	if _, err := s.playerRepository.Creating(playerEntity); err != nil {
		return err
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingReq *_adminModel.AdminCreatingReq) error {
	return nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {
	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return false
	}
	return player != nil
}

func (s *googleOAuth2Service) isThisGuyIsReallyAdmin(adminID string) bool {
	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return false
	}
	return admin != nil
}
