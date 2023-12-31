package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type Transaction struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *Transaction {
	return &Transaction{
		db: db,
	}
}

func (t *Transaction) CreateTransaction(ctx context.Context, transaction *models.Transaction) (uint64, error) {
	var id uint64
	err := t.db.QueryRowContext(
		ctx,
		"INSERT INTO public.transactions (type, from_wallet_id, to_wallet_id, coins, rubles, status, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,NOW(),NOW()) RETURNING id",
		transaction.Type,
		transaction.FromWalletID,
		transaction.ToWalletID,
		transaction.Coins,
		transaction.Rubles,
		transaction.Status,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *Transaction) UpdateStatusByID(ctx context.Context, id uint64, status string) error {
	_, err := t.db.ExecContext(ctx, "UPDATE public.transactions SET (status, updated_at) = ($1, NOW()) WHERE id = $2", status, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) GetTransactionByID(ctx context.Context, id uint64) (*models.Transaction, error) {
	var transaction models.Transaction
	err := t.db.GetContext(ctx, &transaction, "SELECT * FROM public.transactions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *Transaction) GetTransactionByToWalletIDAndFromWalletID(ctx context.Context, fromWalletID, toWalletID uint64) (*models.Transaction, error) {
	var transaction models.Transaction
	err := t.db.GetContext(ctx, &transaction, "SELECT * FROM public.transactions WHERE to_wallet_id = $1 AND from_wallet_id = $2", toWalletID, fromWalletID)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *Transaction) GetNewTransferTransactionsByToWalletID(ctx context.Context, toWalletID uint64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	if err := t.db.SelectContext(ctx, &transactions, "SELECT * FROM public.transactions WHERE to_wallet_id = $1 AND type = $2 AND status = $3", toWalletID, models.TransactionTransferType, models.TransactionNewStatus); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *Transaction) GetTransactionsByWalletID(ctx context.Context, walletID uint64, transactionType, status string) ([]*models.Transaction, error) {
	query := "SELECT * FROM public.transactions WHERE (to_wallet_id = $1 OR from_wallet_id = $2)"
	if transactionType != "" {
		query += fmt.Sprintf(" AND type = %s", transactionType)
	}

	if status != "" {
		query += fmt.Sprintf(" AND status = %s", status)
	}

	query += " ORDER BY id DESC"

	var transactions []*models.Transaction
	if err := t.db.SelectContext(ctx, &transactions, query, walletID, walletID); err != nil {
		return nil, err
	}

	return transactions, nil
}
