package fund

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/fund"
)

func (f *FundEndpoint) GetFundsV1(ctx context.Context, req *fund.GetFundsV1Request) (*fund.GetFundsV1Response, error) {
	funds, err := f.services.FundService.GetFunds(ctx)
	if err != nil {
		return nil, err
	}

	var fundResponse = make([]*fund.GetFundMessage, 0, len(funds))
	for _, fundModel := range funds {
		fundResponse = append(fundResponse, &fund.GetFundMessage{
			Id:          fundModel.ID,
			Name:        fundModel.Name,
			Description: fundModel.Description,
		})
	}

	return &fund.GetFundsV1Response{Funds: fundResponse}, nil
}
