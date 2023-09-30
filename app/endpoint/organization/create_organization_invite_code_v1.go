package organization

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
	"github.com/NameLessCorporation/active-charity-backend/tools/helpers"
)

func (o *OrganizationEndpoint) CreateOrganizationInviteCodeV1(ctx context.Context, req *organization.CreateOrganizationInviteCodeV1Request) (*organization.CreateOrganizationInviteCodeV1Response, error) {
	userID, err := o.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), o.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var user *models.User
	user, err = o.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.OrganizationID == 0 {
		return nil, status.Error(codes.PermissionDenied, "Пользователь не состоит в организации")
	}

	var code = helpers.GenerateRandomString(8)
	if _, err = o.services.InviteCodeService.CreateInviteCode(ctx, &models.InviteCode{
		Code:           code,
		OrganizationID: user.OrganizationID,
		OwnerID:        userID,
	}); err != nil {
		return nil, err
	}

	return &organization.CreateOrganizationInviteCodeV1Response{InviteCode: code}, nil
}
