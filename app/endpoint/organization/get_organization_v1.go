package organization

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
)

func (o *OrganizationEndpoint) GetOrganizationV1(ctx context.Context, req *organization.GetOrganizationV1Request) (*organization.GetOrganizationV1Response, error) {
	userID, err := o.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), o.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var user *models.User
	user, err = o.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var org *models.Organization
	org, err = o.services.OrganizationService.GetOrganizationByID(ctx, user.OrganizationID)
	if err != nil {
		return nil, err
	}

	return &organization.GetOrganizationV1Response{
		Name: org.Name,
	}, nil
}
