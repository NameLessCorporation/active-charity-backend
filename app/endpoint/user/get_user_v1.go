package user

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) GetUserV1(ctx context.Context, req *user.GetUserV1Request) (*user.GetUserV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var usr *models.User
	usr, err = u.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var wallet *models.Wallet
	wallet, err = u.services.WalletService.GetWalletByID(ctx, usr.WalletID)
	if err != nil {
		return nil, err
	}

	return &user.GetUserV1Response{
		Email:          usr.Email,
		Name:           usr.Name,
		DateOfBirthday: usr.DateOfBirthday,
		Coins:          wallet.Coins,
		Rubles:         wallet.Rubles,
	}, nil
}
