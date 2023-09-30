package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type InviteCode struct {
	db *sqlx.DB
}

func NewInviteCodeRepository(db *sqlx.DB) *InviteCode {
	return &InviteCode{
		db: db,
	}
}

func (i *InviteCode) CreateInviteCode(ctx context.Context, code *models.InviteCode) (uint64, error) {
	var id uint64
	err := i.db.QueryRowContext(
		ctx,
		"INSERT INTO public.invite_codes (code, organization_id, owner_id, created_at, updated_at) VALUES ($1,$2, $3, NOW(), NOW()) RETURNING id",
		code.Code,
		code.OrganizationID,
		code.OwnerID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (i *InviteCode) GetInviteCodeByCode(ctx context.Context, code string) (*models.InviteCode, error) {
	var inviteCode models.InviteCode
	err := i.db.GetContext(ctx, &inviteCode, "SELECT * FROM public.invite_codes WHERE code = $1", code)
	if err != nil {
		return nil, err
	}

	return &inviteCode, nil
}
