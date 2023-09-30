package models

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	ID             uint64    `db:"id"`
	Email          string    `db:"email"`
	Name           string    `db:"name"`
	Password       string    `db:"password"`
	DateOfBirthday string    `db:"date_of_birthday"`
	OrganizationID uint64    `db:"organization_id"`
	WalletID       uint64    `db:"wallet_id"`
	FundID         uint64    `db:"fund_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (c *User) Validate() error {
	if len(c.Password) > 64 {
		return status.Error(codes.InvalidArgument, "Длина пароля больше 64 символов")
	}
	if len(c.Password) < 8 {
		return status.Error(codes.InvalidArgument, "Длина пароля меньше 8 символов")
	}

	return nil
}
