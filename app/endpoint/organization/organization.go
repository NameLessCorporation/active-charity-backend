package organization

import (
	"github.com/NameLessCorporation/active-charity-backend/app/balance_operations"
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type OrganizationEndpoint struct {
	services   *service.Services
	config     *config.Config
	operations balance_operations.BalanceOperations
}

func NewOrganizationEndpoint(services *service.Services, config *config.Config, operations balance_operations.BalanceOperations) *OrganizationEndpoint {
	return &OrganizationEndpoint{
		config:     config,
		services:   services,
		operations: operations,
	}
}
