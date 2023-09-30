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
		"INSERT INTO public.organizations (name, owner_id, created_at, updated_at) VALUES ($1,$2, NOW(), NOW()) RETURNING id",
		organization.Name,
		organization.OwnerID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
