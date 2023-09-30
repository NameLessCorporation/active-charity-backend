package models

import (
	"time"
)

type Token struct {
	ID           uint64    `db:"id"`
	AccessToken  string    `db:"access_token"`
	RefreshToken string    `db:"refresh_token"`
	Exp          time.Time `db:"exp"`
	UserID       uint64    `db:"user_id"`
}
