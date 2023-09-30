package models

import "time"

const (
	WalletOrganizationType = "organization"
	WalletUserType         = "user"
)

type Wallet struct {
	ID        uint64    `db:"id"`
	Type      string    `db:"type"`
	Coins     uint64    `db:"coins"`
	Rubles    uint64    `db:"rubles"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
