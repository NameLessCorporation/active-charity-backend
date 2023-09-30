package user

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) JoinUserToOrganizationV1(ctx context.Context, req *user.JoinUserToOrganizationV1Request) (*user.JoinUserToOrganizationV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var inviteCode *models.InviteCode
	inviteCode, err = u.services.InviteCodeService.GetInviteCodeByCode(ctx, req.GetOrganizationInviteCode())
	if err != nil {
		return nil, err
	}

	if err = u.services.UserService.UpdateOrganizationIDByID(ctx, userID, inviteCode.OrganizationID); err != nil {
		return nil, err
	}

	return &user.JoinUserToOrganizationV1Response{}, nil
}
