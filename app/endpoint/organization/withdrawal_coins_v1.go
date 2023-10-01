package organization

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
)

func (o *OrganizationEndpoint) WithdrawalCoinsV1(ctx context.Context, req *organization.WithdrawalCoinsV1Request) (*organization.WithdrawalCoinsV1Response, error) {
	userID, err := o.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), o.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var usr *models.User
	usr, err = o.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var org *models.Organization
	org, err = o.services.OrganizationService.GetOrganizationByID(ctx, usr.OrganizationID)

	if org.OwnerID != userID {
		return nil, status.Error(codes.PermissionDenied, "Пользователь не владеет организациец")
	}

	var transactionID uint64
	transactionID, err = o.operations.WithdrawalCoinsFromOrganization(ctx, req.GetCoins(), org.WalletID)
	if err != nil {
		return nil, err
	}

	return &organization.WithdrawalCoinsV1Response{TransactionId: transactionID}, nil
}
