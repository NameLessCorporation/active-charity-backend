package models

import "time"

type Organization struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	OwnerID   uint64    `db:"owner_id"`
	WalletID  uint64    `db:"wallet_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
