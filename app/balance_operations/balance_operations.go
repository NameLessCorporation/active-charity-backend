package balance_operations

import "github.com/NameLessCorporation/active-charity-backend/app/service"

type BalanceOperations interface {
}

type balanceOperations struct {
	services service.Services
}

func NewBalanceOperations(services service.Services) BalanceOperations {
	return &balanceOperations{
		services: services,
	}
}

func (b *balanceOperations) ConvertCoinsToRubles(coins uint64) uint64 {
	return
}
