package fund

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/fund"
)

func (f *FundEndpoint) CreateFundV1(ctx context.Context, req *fund.CreateFundV1Request) (*fund.CreateFundV1Response, error) {
	userID, err := f.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), f.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var fundID uint64
	fundID, err = f.services.FundService.CreateFund(ctx, &models.Fund{
		Name:        req.GetFund().GetName(),
		Description: req.GetFund().GetDescription(),
		OwnerID:     userID,
	})
	if err != nil {
		return nil, err
	}

	if err = f.services.UserService.UpdateFundIDByID(ctx, userID, fundID); err != nil {
		return nil, err
	}

	return &fund.CreateFundV1Response{}, nil
}
