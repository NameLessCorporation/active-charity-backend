package endpoint

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/auth"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

type EndpointContainer struct {
	AuthService AuthServiceInter
	UserService UserServiceInter
}

func NewEndpointContainer(auth AuthServiceInter, user UserServiceInter) *EndpointContainer {
	return &EndpointContainer{
		AuthService: auth,
		UserService: user,
	}
}

type AuthServiceInter interface {
	AuthWithCredentials(ctx context.Context, req *auth.AuthWithCredentialsRequest) (*auth.AccessRefreshTokens, error)
	RefreshTokens(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error)
	AuthWithAccessToken(ctx context.Context, req *auth.AuthWithAccessTokenRequest) (*auth.AuthWithAccessTokenResponse, error)
}

type UserServiceInter interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserMessageResponse, error)
}
