package organization

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
)

func (o *OrganizationEndpoint) GetOrganizationUsersV1(ctx context.Context, req *organization.GetOrganizationUsersV1Request) (*organization.GetOrganizationUsersV1Response, error) {
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
		return nil, status.Error(codes.PermissionDenied, "Пользователь не владеет организацией")
	}

	var users []*models.User
	users, err = o.services.UserService.GetUsersByOrganizationID(ctx, usr.OrganizationID)
	if err != nil {
		return nil, err
	}

	var usersResponse = make([]*organization.User, 0, len(users))
	for _, userModel := range users {
		var fundName string
		if userModel.FundID != 0 {
			var fund *models.Fund
			fund, err = o.services.FundService.GetFundByID(ctx, userModel.FundID)
			if err != nil {
				return nil, err
			}

			fundName = fund.Name
		}

		var wallet *models.Wallet
		wallet, err = o.services.WalletService.GetWalletByID(ctx, userModel.WalletID)
		if err != nil {
			return nil, err
		}

		usersResponse = append(usersResponse, &organization.User{
			Id:             userID,
			Email:          userModel.Email,
			Name:           userModel.Name,
			FundName:       fundName,
			Coins:          wallet.Coins,
			Rubles:         wallet.Rubles,
			DateOfBirthday: userModel.DateOfBirthday,
			CreatedAt:      userModel.CreatedAt.String(),
		})
	}

	return &organization.GetOrganizationUsersV1Response{
		Users: usersResponse,
	}, nil
}
