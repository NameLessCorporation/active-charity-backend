package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type Fund struct {
	db *sqlx.DB
}

func NewFundRepository(db *sqlx.DB) *Fund {
	return &Fund{
		db: db,
	}
}

func (f *Fund) CreateFund(ctx context.Context, fund *models.Fund) (uint64, error) {
	var id uint64
	err := f.db.QueryRowContext(
		ctx,
		"INSERT INTO public.funds (name, description, owner_id, created_at, updated_at) VALUES ($1,$2,$3,NOW(),NOW()) RETURNING id",
		fund.Name,
		fund.Description,
		fund.OwnerID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (f *Fund) GetFundByID(ctx context.Context, id uint64) (*models.Fund, error) {
	var fund models.Fund
	err := f.db.GetContext(ctx, &fund, "SELECT * FROM public.funds WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &fund, nil
}

func (f *Fund) GetFunds(ctx context.Context) ([]*models.Fund, error) {
	var funds []*models.Fund
	if err := f.db.SelectContext(ctx, &funds, "SELECT * FROM public.funds"); err != nil {
		return nil, err
	}

	return funds, nil
}
