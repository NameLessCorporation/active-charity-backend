package user

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) SelectUserFundV1(ctx context.Context, req *user.SelectUserFundV1Request) (*user.SelectUserFundV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if err = u.services.UserService.UpdateFundIDByID(ctx, userID, req.GetFundId()); err != nil {
		return nil, err
	}

	return &user.SelectUserFundV1Response{}, nil
}
