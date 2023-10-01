package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type User struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) GetWalletIdById(ctx context.Context, id uint64) (uint64, error) {
	var walletId uint64
	err := u.db.GetContext(ctx, &walletId, "SELECT wallet_id FROM users WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("no rows")
	}

	return walletId, nil
}

func (u *User) IsExistByEmail(ctx context.Context, email string) bool {
	var id uint64
	err := u.db.GetContext(ctx, &id, "SELECT id FROM public.users WHERE email = $1", email)
	if err != nil {
		return false
	}

	if id == 0 {
		return false
	}

	return true
}

func (u *User) GetIDByCredentials(ctx context.Context, credentials *models.Credential) (uint64, error) {
	var id uint64
	err := u.db.GetContext(
		ctx,
		&id,
		"SELECT id FROM public.users WHERE email = $1 AND password = $2",
		credentials.Email,
		credentials.Password,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, "SELECT * FROM public.users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	var id uint64
	err := u.db.QueryRowContext(
		ctx,
		"INSERT INTO public.users (email, password, name, date_of_birthday, organization_id, wallet_id, fund_id, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,NOW(),NOW()) RETURNING id",
		user.Email,
		user.Password,
		user.Name,
		user.DateOfBirthday,
		user.OrganizationID,
		user.WalletID,
		user.FundID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) GetIDByEmail(ctx context.Context, email string) (uint64, error) {
	var id uint64
	err := u.db.GetContext(ctx, &id, "SELECT id FROM public.users WHERE email = $1", email)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) UpdateOrganizationIDByID(ctx context.Context, id, organizationID uint64) error {
	_, err := u.db.ExecContext(ctx, "UPDATE public.users SET (organization_id, updated_at) = ($1, NOW()) WHERE id = $2", organizationID, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateFundIDByID(ctx context.Context, id, fundID uint64) error {
	_, err := u.db.ExecContext(ctx, "UPDATE public.users SET (fund_id, updated_at) = ($1, NOW()) WHERE id = $2", fundID, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByWalletID(ctx context.Context, walletID uint64) (*models.User, error) {
	var user models.User
	err := u.db.GetContext(ctx, &user, "SELECT * FROM public.users WHERE wallet_id = $1", walletID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetUsersByOrganizationID(ctx context.Context, id uint64) ([]*models.User, error) {
	var users []*models.User

	err := u.db.SelectContext(ctx, &users, "select * from users where organization_id = $1", id)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) GetOrganizationByUserID(ctx context.Context, id uint64) (uint64, error) {
	var organizationID uint64

	err := u.db.GetContext(ctx, &organizationID, "select organization_id from users where id = $1", id)
	if err != nil {
		return 0, err
	}

	return organizationID, nil
}
