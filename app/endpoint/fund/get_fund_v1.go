package fund

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/fund"
)

func (f *FundEndpoint) GetFundV1(ctx context.Context, req *fund.GetFundV1Request) (*fund.GetFundV1Response, error) {
	userID, err := f.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), f.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var user *models.User
	user, err = f.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var fundModel *models.Fund
	fundModel, err = f.services.FundService.GetFundByID(ctx, user.FundID)
	if err != nil {
		return nil, err
	}

	return &fund.GetFundV1Response{
		Fund: &fund.GetFundMessage{
			Id:          fundModel.ID,
			Name:        fundModel.Name,
			Description: fundModel.Description,
		},
	}, nil
}
