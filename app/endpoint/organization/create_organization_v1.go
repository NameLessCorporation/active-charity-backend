package organization

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
)

func (o *OrganizationEndpoint) CreateOrganizationV1(
	ctx context.Context,
	req *organization.CreateOrganizationV1Request,
) (*organization.CreateOrganizationV1Response, error) {
	userID, err := o.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), o.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var id uint64
	id, err = o.services.OrganizationService.CreateOrganization(ctx, &models.Organization{
		Name:    req.GetName(),
		OwnerID: userID,
	})
	if err != nil {
		return nil, err
	}

	if err = o.services.UserService.UpdateOrganizationIDByID(ctx, userID, id); err != nil {
		return nil, err
	}

	return &organization.CreateOrganizationV1Response{}, nil
}
