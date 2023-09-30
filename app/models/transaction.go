package models

import "time"

const (
	TransactionTransferType   = "transfer"
	TransactionWithdrawalType = "withdrawal"
	TransactionInsertionType  = "insertion"

	TransactionNewStatus      = "new"
	TransactionApprovedStatus = "approved"
	TransactionRejectedStatus = "rejected"
)

type Transaction struct {
	ID           uint64    `db:"id"`
	Type         string    `db:"type"`
	FromWalletID uint64    `db:"from_wallet_id"`
	ToWalletID   uint64    `db:"to_wallet_id"`
	Coins        uint64    `db:"coins"`
	Rubles       uint64    `db:"rubles"`
	Status       string    `db:"status"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
