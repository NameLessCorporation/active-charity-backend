package auth

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/extra/auth"
	jwt_helper "github.com/NameLessCorporation/active-charity-backend/tools/jwt"
)

func (a *AuthEndpoint) AuthWithAccessTokenV1(_ context.Context, req *auth.AuthWithAccessTokenV1Request) (*auth.AuthWithAccessTokenV1Response, error) {
	token, err := jwt_helper.ParseJWT([]byte(a.config.JWT.SecretKey), req.AccessToken)
	if err != nil {
		return nil, err
	}

	if token.ExpiresAt < time.Now().Unix() {
		return nil, status.Error(codes.Unauthenticated, "Срок действия токена доступа истек")
	}

	return &auth.AuthWithAccessTokenV1Response{}, nil
}
