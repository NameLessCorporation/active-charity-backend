package service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/tools/helpers"
	jwt_helper "github.com/NameLessCorporation/active-charity-backend/tools/jwt"
)

func (s *Service) CreateToken(ctx context.Context, token *models.Token) error {
	if err := s.repository.TokenRepository.CreateToken(ctx, token); err != nil {
		return status.Error(codes.Internal, "Ошибка создания токена")
	}

	return nil
}

func (s *Service) GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error) {
	token, err := s.repository.TokenRepository.GetTokenByAccessToken(ctx, accessToken)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения токена")
	}

	return token, nil
}

func (s *Service) DeleteTokenByAccessToken(ctx context.Context, accessToken string) error {
	if err := s.repository.TokenRepository.DeleteTokenByAccessToken(ctx, accessToken); err != nil {
		return status.Error(codes.Internal, "Ошибка удаления токена доступа")
	}

	return nil
}

func (s *Service) GenerateTokens(ctx context.Context, userID uint64, jwtSecretKey string, expirationTime time.Duration) (*models.Token, error) {
	refreshToken := helpers.GenerateRandomString(64)
	accessToken, err := jwt_helper.CreateJWT([]byte(jwtSecretKey), &jwt_helper.Claims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка генерации токена")
	}

	token := &models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
		Exp:          time.Now().Add(expirationTime),
	}

	return token, nil
}

func (s *Service) GetUserIDByAccessToken(ctx context.Context, accessToken string, secretKey string) (uint64, error) {
	if accessToken == "" {
		return 0, status.Error(codes.Unauthenticated, "Данного токена не существует")
	}

	claims, err := jwt_helper.ParseJWT([]byte(secretKey), accessToken)
	if err != nil {
		return 0, status.Error(codes.Unauthenticated, "Токен не действителен")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return 0, status.Error(codes.Unauthenticated, "Время действия токена истекло")
	}

	if err = claims.Valid(); err != nil {
		return 0, status.Error(codes.Internal, "Ошибка валидации токена")
	}

	return claims.ID, nil
}
