package activity

import (
	"github.com/NameLessCorporation/active-charity-backend/app/balance_operations"
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type ActivityEndpoint struct {
	services   *service.Services
	config     *config.Config
	operations balance_operations.BalanceOperations
}

func NewActivityEndpoint(services *service.Services, config *config.Config, operations balance_operations.BalanceOperations) *ActivityEndpoint {
	return &ActivityEndpoint{
		config:     config,
		services:   services,
		operations: operations,
	}
}
