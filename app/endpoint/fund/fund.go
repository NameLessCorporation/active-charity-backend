package fund

import (
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type FundEndpoint struct {
	services *service.Services
	config   *config.Config
}

func NewFundEndpoint(services *service.Services, config *config.Config) *FundEndpoint {
	return &FundEndpoint{
		config:   config,
		services: services,
	}
}
