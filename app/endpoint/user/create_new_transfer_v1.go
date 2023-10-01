package user

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) CreateNewTransferV1(ctx context.Context, req *user.CreateNewTransferV1Request) (*user.CreateNewTransferV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var usr *models.User
	usr, err = u.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var org *models.Organization
	org, err = u.services.OrganizationService.GetOrganizationByID(ctx, usr.OrganizationID)
	if err != nil {
		return nil, err
	}

	var transactionID uint64
	transactionID, err = u.operations.CreateTransferCoinsToOrganization(ctx, req.GetCoins(), usr.WalletID, org.WalletID)
	if err != nil {
		return nil, err
	}

	return &user.CreateNewTransferV1Response{TransactionId: transactionID}, nil
}
