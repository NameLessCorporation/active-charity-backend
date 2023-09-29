package service

import (
	"context"
	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/app/repository"
	"github.com/NameLessCorporation/active-charity-backend/config"
	"time"
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

type IConfigService interface {
}

type IRoleService interface {
	IsExistByRole(ctx context.Context, role string) bool
	GetRoleIDByName(ctx context.Context, name string) (*models.Role, error)
	GetRoleByRoleID(ctx context.Context, roleID uint64) (*models.Role, error)
	IsRoleAllowedForUser(ctx context.Context, userID uint64, role string) (bool, error)
	CreateRole(ctx context.Context, role *models.Role) (uint64, error)
}

type IValueService interface {
}

type Services struct {
	TokenService  ITokenService
	UserService   IUserService
	ConfigService IConfigService
	RoleService   IRoleService
	ValueService  IValueService
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
		TokenService:  s,
		UserService:   s,
		ConfigService: s,
		RoleService:   s,
		ValueService:  s,
	}
}
