package endpoint

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/auth"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

type EndpointContainer struct {
	AuthService         AuthServiceInter
	UserService         UserServiceInter
	OrganizationService OrganizationServiceInter
}

func NewEndpointContainer(
	auth AuthServiceInter,
	user UserServiceInter,
	organizationService OrganizationServiceInter,
) *EndpointContainer {
	return &EndpointContainer{
		AuthService:         auth,
		UserService:         user,
		OrganizationService: organizationService,
	}
}

type AuthServiceInter interface {
	AuthWithCredentialsV1(ctx context.Context, req *auth.AuthWithCredentialsV1Request) (*auth.AccessRefreshTokens, error)
	RefreshTokensV1(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error)
	AuthWithAccessTokenV1(ctx context.Context, req *auth.AuthWithAccessTokenV1Request) (*auth.AuthWithAccessTokenV1Response, error)
}

type UserServiceInter interface {
	CreateUserV1(ctx context.Context, req *user.CreateUserV1Request) (*user.CreateUserV1Response, error)
	JoinUserToOrganizationV1(ctx context.Context, req *user.JoinUserToOrganizationV1Request) (*user.JoinUserToOrganizationV1Response, error)
}

type OrganizationServiceInter interface {
	CreateOrganizationV1(ctx context.Context, req *organization.CreateOrganizationV1Request) (*organization.CreateOrganizationV1Response, error)
	CreateOrganizationInviteCodeV1(ctx context.Context, req *organization.CreateOrganizationInviteCodeV1Request) (*organization.CreateOrganizationInviteCodeV1Response, error)
}
