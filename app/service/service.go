package service

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/app/repository"
	"github.com/NameLessCorporation/active-charity-backend/config"
)

type IActivityService interface {
	TrackSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) error
	GetActivityList(ctx context.Context) ([]*models.Activity, error)
}

type ITokenService interface {
	CreateToken(ctx context.Context, token *models.Token) error
	GenerateTokens(ctx context.Context, userID uint64, jwtSecretKey string, expirationTime time.Duration) (*models.Token, error)
	GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error)
	DeleteTokenByAccessToken(ctx context.Context, accessToken string) error
	GetUserIDByAccessToken(ctx context.Context, accessToken string, secretKey string) (uint64, error)
}

type IUserService interface {
	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	IsExistByEmail(ctx context.Context, email string) bool
	GetIDByCredentials(ctx context.Context, credentials *models.Credential) (uint64, error)
	GetUserByID(ctx context.Context, id uint64) (*models.User, error)
	GetIDByEmail(ctx context.Context, email string) (uint64, error)
	UpdateOrganizationIDByID(ctx context.Context, id, organizationID uint64) error
	UpdateFundIDByID(ctx context.Context, id, fundID uint64) error
}

type IOrganizationService interface {
	CreateOrganization(ctx context.Context, organization *models.Organization) (uint64, error)
	GetOrganizationByID(ctx context.Context, id uint64) (*models.Organization, error)
}

type IInviteCodeService interface {
	CreateInviteCode(ctx context.Context, inviteCode *models.InviteCode) (uint64, error)
	GetInviteCodeByCode(ctx context.Context, code string) (*models.InviteCode, error)
}

type IWalletService interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) (uint64, error)
	GetWalletByID(ctx context.Context, id uint64) (*models.Wallet, error)
	UpdateCoinsByID(ctx context.Context, id, coins uint64) error
	UpdateRublesByID(ctx context.Context, id, rubles uint64) error
}

type ITransactionService interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (uint64, error)
	UpdateStatusByID(ctx context.Context, id uint64, status string) error
	GetTransactionByID(ctx context.Context, id uint64) (*models.Transaction, error)
	GetTransactionByToWalletIDAndFromWalletID(ctx context.Context, fromWalletID, toWalletID uint64) (*models.Transaction, error)
}

type IFundService interface {
	CreateFund(ctx context.Context, fund *models.Fund) (uint64, error)
	GetFundByID(ctx context.Context, id uint64) (*models.Fund, error)
	GetFunds(ctx context.Context) ([]*models.Fund, error)
}

type Services struct {
	TokenService        ITokenService
	UserService         IUserService
	OrganizationService IOrganizationService
	InviteCodeService   IInviteCodeService
	ActivityService     IActivityService
	WalletService       IWalletService
	TransactionService  ITransactionService
	FundService         IFundService
}

type Service struct {
	repository *repository.Repository
	config     *config.Config
	Services   *Services
	logger     *zap.Logger
}

func NewService(repository *repository.Repository, config *config.Config, logger *zap.Logger) *Service {
	return &Service{
		repository: repository,
		config:     config,
		logger:     logger,
	}
}

func (s *Service) InitServices() {
	s.Services = &Services{
		TokenService:        s,
		UserService:         s,
		OrganizationService: s,
		InviteCodeService:   s,
		WalletService:       s,
		TransactionService:  s,
		ActivityService:     s,
		FundService:         s,
	}
}
