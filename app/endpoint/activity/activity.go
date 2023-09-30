package activity

import (
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type ActivityEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewActivityEndpoint(services *service.Services, config *config.Config) *ActivityEndpoint {
	return &ActivityEndpoint{
		config:   config,
		services: services,
	}
}
