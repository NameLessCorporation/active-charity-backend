package organization

import (
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type OrganizationEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewOrganizationEndpoint(services *service.Services, config *config.Config) *OrganizationEndpoint {
	return &OrganizationEndpoint{
		config:   config,
		services: services,
	}
}
