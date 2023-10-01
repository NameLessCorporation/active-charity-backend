package user

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) GetTopV1(ctx context.Context, req *user.GetTopV1Request) (*user.GetTopV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	organizationID, err := u.services.UserService.GetOrganizationByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	users, err := u.services.UserService.GetUsersByOrganizationID(ctx, organizationID)
	if err != nil {
		return nil, err
	}
}
