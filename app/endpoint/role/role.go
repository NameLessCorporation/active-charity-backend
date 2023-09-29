package role

import (
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type RoleEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewRoleEndpoint(services *service.Services, config *config.Config) *RoleEndpoint {
	return &RoleEndpoint{
		config:   config,
		services: services,
	}
}
