package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type ActivityRepository interface {
	TrackCurrentPeriodSteps(ctx context.Context, steps uint32, periodId uint64) error
	TrackNewPeriodSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) error
	GetActivityList(ctx context.Context) ([]*models.Activity, error)
	IsActiveStepsPeriod(ctx context.Context, userId uint64) (bool, error)
	GetCurrentPeriodId(ctx context.Context, userId uint64) (uint64, uint32, error)
	TrackPullUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error
	TrackPushUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error
	TrackBenchPress(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error
	TrackCycling(ctx context.Context, metres uint32, activityId uint64, userId uint64) error
	TrackCrunches(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error

	GetStepsValue(ctx context.Context, userID uint64) (sql.NullInt32, error)
	GetBenchPressValue(ctx context.Context, userID uint64) (sql.NullInt32, error)
	GetCrunchesValue(ctx context.Context, userID uint64) (sql.NullInt32, error)
	GetCyclingValue(ctx context.Context, userID uint64) (sql.NullInt32, error)
	GetPullUpValue(ctx context.Context, userID uint64) (sql.NullInt32, error)
	GetPushUpValue(ctx context.Context, userID uint64) (sql.NullInt32, error)

	GetUserFavouriteActivity(ctx context.Context, userID uint64) (string, error)
	GetUserMostEarnedActivity(ctx context.Context, userID uint64) (string, error)

	GetUserActivityAnalytics(ctx context.Context, userID uint64) ([]models.ActivityAnalytics, error)
	GetStepsGraph(ctx context.Context, userID uint64) ([]*models.Graph, error)
	GetPushUpGraph(ctx context.Context, userID uint64) ([]*models.Graph, error)
	GetCrunchesGraph(ctx context.Context, userID uint64) ([]*models.Graph, error)
	GetCycleGraph(ctx context.Context, userID uint64) ([]*models.Graph, error)
	GetPullUpGraph(ctx context.Context, userID uint64) ([]*models.Graph, error)
	GetBenchPressGraph(ctx context.Context, userID uint64) ([]*models.Graph, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	IsExistByEmail(ctx context.Context, email string) bool
	GetIDByCredentials(ctx context.Context, credentials *models.Credential) (uint64, error)
	GetUserByID(ctx context.Context, id uint64) (*models.User, error)
	GetIDByEmail(ctx context.Context, email string) (uint64, error)
	UpdateOrganizationIDByID(ctx context.Context, id, organizationID uint64) error
	UpdateFundIDByID(ctx context.Context, id, fundID uint64) error
	GetUserByWalletID(ctx context.Context, walletID uint64) (*models.User, error)
	GetWalletIdById(ctx context.Context, id uint64) (uint64, error)
	GetUsersByOrganizationID(ctx context.Context, id uint64) ([]*models.User, error)
	GetOrganizationByUserID(ctx context.Context, id uint64) (uint64, error)
}

type TokenRepository interface {
	CreateToken(ctx context.Context, token *models.Token) error
	GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error)
	DeleteTokenByAccessToken(ctx context.Context, accessToken string) error
}

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, organization *models.Organization) (uint64, error)
	GetOrganizationByID(ctx context.Context, id uint64) (*models.Organization, error)
	GetOrganizationByWalletID(ctx context.Context, walletID uint64) (*models.Organization, error)
}

type InviteCodeRepository interface {
	CreateInviteCode(ctx context.Context, code *models.InviteCode) (uint64, error)
	GetInviteCodeByCode(ctx context.Context, code string) (*models.InviteCode, error)
}

type WalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) (uint64, error)
	GetWalletByID(ctx context.Context, id uint64) (*models.Wallet, error)
	UpdateCoinsByID(ctx context.Context, id, coins uint64) error
	UpdateRublesByID(ctx context.Context, id, rubles uint64) error
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (uint64, error)
	UpdateStatusByID(ctx context.Context, id uint64, status string) error
	GetTransactionByID(ctx context.Context, id uint64) (*models.Transaction, error)
	GetTransactionByToWalletIDAndFromWalletID(ctx context.Context, fromWalletID, toWalletID uint64) (*models.Transaction, error)
	GetNewTransferTransactionsByToWalletID(ctx context.Context, toWalletID uint64) ([]*models.Transaction, error)
	GetTransactionsByWalletID(ctx context.Context, walletID uint64, transactionType, status string) ([]*models.Transaction, error)
}

type FundRepository interface {
	CreateFund(ctx context.Context, fund *models.Fund) (uint64, error)
	GetFundByID(ctx context.Context, id uint64) (*models.Fund, error)
	GetFunds(ctx context.Context) ([]*models.Fund, error)
}

type Repository struct {
	UserRepository         UserRepository
	TokenRepository        TokenRepository
	OrganizationRepository OrganizationRepository
	InviteCodeRepository   InviteCodeRepository
	ActivityRepository     ActivityRepository
	WalletRepository       WalletRepository
	TransactionRepository  TransactionRepository
	FundRepository         FundRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:         NewUserRepository(db),
		TokenRepository:        NewTokenRepository(db),
		OrganizationRepository: NewOrganizationRepository(db),
		InviteCodeRepository:   NewInviteCodeRepository(db),
		ActivityRepository:     NewActivityRepository(db),
		WalletRepository:       NewWalletRepository(db),
		TransactionRepository:  NewTransactionRepository(db),
		FundRepository:         NewFundRepository(db),
	}
}
