package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/auth"
	"github.com/NameLessCorporation/active-charity-backend/tools/helpers"
)

func (a *AuthEndpoint) AuthWithCredentialsV1(ctx context.Context, req *auth.AuthWithCredentialsV1Request) (*auth.AccessRefreshTokens, error) {
	cred := &models.Credential{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	if err := cred.Validate(); err != nil {
		return nil, err
	}

	if !a.services.UserService.IsExistByEmail(ctx, req.GetEmail()) {
		return nil, status.Error(codes.AlreadyExists, "Данный email уже зарегистрирован")
	}

	hashPassword, err := helpers.NewMD5Hash(req.GetPassword())
	if err != nil {
		return nil, err
	}

	cred.Password = hashPassword

	var id uint64
	id, err = a.services.UserService.GetIDByCredentials(ctx, cred)
	if err != nil {
		return nil, err
	}

	var token *models.Token
	token, err = a.services.TokenService.GenerateTokens(ctx, id, a.config.JWT.SecretKey, a.config.JWT.ExpirationTime)
	if err != nil {
		return nil, err
	}

	if err = a.services.TokenService.CreateToken(ctx, token); err != nil {
		return nil, status.Error(codes.Internal, "Ошибка создания пары токенов")
	}

	return &auth.AccessRefreshTokens{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
