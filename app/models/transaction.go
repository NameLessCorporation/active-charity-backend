package models

const (
	TransactionTransferType   = "transfer"
	TransactionWithdrawalType = "withdrawal"
	TransactionInsertionType  = "insertion"

	TransactionNewStatus          = "new"
	TransactionInProcessingStatus = "in processing"
	TransactionApprovedStatus     = "approved"
	TransactionRejectedStatus     = "rejected"
)

type Transaction struct {
	ID           uint64 `db:"id"`
	Type         string `db:"type"`
	FromWalletID string `db:"from_wallet_id"`
	ToWalletID   string `db:"to_wallet_id"`
	Coins        uint64 `db:"coins"`
	Rubles       uint64 `db:"rubles"`
	Status       string `db:"status"`
}
