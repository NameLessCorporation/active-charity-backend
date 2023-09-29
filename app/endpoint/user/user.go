package user

import (
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type UserEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewUserEndpoint(services *service.Services, config *config.Config) *UserEndpoint {
	return &UserEndpoint{
		config:   config,
		services: services,
	}
}
