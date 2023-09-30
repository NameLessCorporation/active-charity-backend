package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (uint64, error)
	IsExistByEmail(ctx context.Context, email string) bool
	GetIDByCredentials(ctx context.Context, credentials *models.Credential) (uint64, error)
	GetUserByID(ctx context.Context, id uint64) (*models.User, error)
	GetIDByEmail(ctx context.Context, email string) (uint64, error)
	UpdateOrganizationIDByID(ctx context.Context, id, organizationID uint64) error
}

type TokenRepository interface {
	CreateToken(ctx context.Context, token *models.Token) error
	GetTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error)
	DeleteTokenByAccessToken(ctx context.Context, accessToken string) error
}

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, organization *models.Organization) (uint64, error)
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
	GetTransactionByToWalletID(ctx context.Context, toWalletID uint64) (*models.Transaction, error)
	GetTransactionByFromWalletID(ctx context.Context, fromWalletID uint64) (*models.Transaction, error)
}

type Repository struct {
	UserRepository         UserRepository
	TokenRepository        TokenRepository
	OrganizationRepository OrganizationRepository
	InviteCodeRepository   InviteCodeRepository
	WalletRepository       WalletRepository
	TransactionRepository  TransactionRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:         NewUserRepository(db),
		TokenRepository:        NewTokenRepository(db),
		OrganizationRepository: NewOrganizationRepository(db),
		InviteCodeRepository:   NewInviteCodeRepository(db),
		WalletRepository:       NewWalletRepository(db),
		TransactionRepository:  NewTransactionRepository(db),
	}
}
