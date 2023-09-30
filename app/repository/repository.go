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

type Repository struct {
	UserRepository         UserRepository
	TokenRepository        TokenRepository
	OrganizationRepository OrganizationRepository
	InviteCodeRepository   InviteCodeRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:         NewUserRepository(db),
		TokenRepository:        NewTokenRepository(db),
		OrganizationRepository: NewOrganizationRepository(db),
		InviteCodeRepository:   NewInviteCodeRepository(db),
	}
}
