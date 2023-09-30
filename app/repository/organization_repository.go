package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type Organization struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) *Organization {
	return &Organization{
		db: db,
	}
}

func (o *Organization) CreateOrganization(ctx context.Context, organization *models.Organization) (uint64, error) {
	var id uint64
	err := o.db.QueryRowContext(
		ctx,
		"INSERT INTO public.organizations (name, owner_id, wallet_id, created_at, updated_at) VALUES ($1,$2,$3,NOW(),NOW()) RETURNING id",
		organization.Name,
		organization.OwnerID,
		organization.WalletID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (o *Organization) GetOrganizationByID(ctx context.Context, id uint64) (*models.Organization, error) {
	var organization models.Organization
	err := o.db.GetContext(ctx, &organization, "SELECT * FROM public.organizations WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}
