package auth

import (
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type AuthEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewAuthEndpoint(services *service.Services, config *config.Config) *AuthEndpoint {
	return &AuthEndpoint{
		config:   config,
		services: services,
	}
}
