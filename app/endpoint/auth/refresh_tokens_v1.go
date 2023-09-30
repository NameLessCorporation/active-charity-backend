package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/auth"
)

func (a *AuthEndpoint) RefreshTokensV1(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error) {
	token, err := a.services.TokenService.GetTokenByAccessToken(ctx, req.GetAccessToken())
	if err != nil {
		return nil, err
	}

	if !(token.RefreshToken == req.GetRefreshToken() && token.AccessToken == req.GetAccessToken()) {
		return nil, status.Error(codes.Unauthenticated, "Некорректный токен")
	}

	var tokenAuth *models.Token
	tokenAuth, err = a.services.TokenService.GenerateTokens(ctx, token.UserID, a.config.JWT.SecretKey, a.config.JWT.ExpirationTime)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка генерации пары токенов")
	}

	err = a.services.TokenService.DeleteTokenByAccessToken(ctx, req.GetAccessToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка удаления пары токенов")
	}

	err = a.services.TokenService.CreateToken(ctx, tokenAuth)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка создания пары токенов")
	}

	return &auth.AccessRefreshTokens{
		AccessToken:  tokenAuth.AccessToken,
		RefreshToken: tokenAuth.RefreshToken,
	}, nil
}
