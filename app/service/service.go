package service

import (
	"context"
	"time"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/app/repository"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type ITokenService interface {
	CreateToken(ctx context.Context, token *models.Token) error
	GenerateTokens(ctx context.Context, userID uint64, jwtSecretKey string, expirationTime time.Duration) (*models.Token, error)
	GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error)
	DeleteTokenByAccessToken(ctx context.Context, accessToken string) error
	GetUserIDByAccessToken(ctx context.Context, accessToken string, secretKey string) (uint64, error)
}

type IUserService interface {
	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	IsExistByLogin(ctx context.Context, login string) bool
	GetUserByCredentials(ctx context.Context, credentials *models.Credentials) (uint64, error)
	GetUserByUserID(ctx context.Context, userID uint64) (*models.User, error)
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	DeleteUserByUserID(ctx context.Context, userID uint64) error
	UpdateUserPasswordAndRoleByUserID(ctx context.Context, userID uint64, password string, role string) error
	GetUserIDByLogin(ctx context.Context, login string) (uint64, error)
}

type Services struct {
	TokenService ITokenService
	UserService  IUserService
}

type Service struct {
	repository *repository.Repository
	config     *config.Config
	Services   *Services
}

func NewService(repository *repository.Repository, config *config.Config) *Service {
	return &Service{
		repository: repository,
		config:     config,
	}
}

func (s *Service) InitServices() {
	s.Services = &Services{
		TokenService: s,
		UserService:  s,
	}
}
