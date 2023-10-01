package user

import (
	"github.com/NameLessCorporation/active-charity-backend/app/balance_operations"
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type UserEndpoint struct {
	services   *service.Services
	config     *config.Config
	operations balance_operations.BalanceOperations
}

func NewUserEndpoint(services *service.Services, config *config.Config, operations balance_operations.BalanceOperations) *UserEndpoint {
	return &UserEndpoint{
		config:     config,
		services:   services,
		operations: operations,
	}
}
