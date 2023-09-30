package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

type Wallet struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *Wallet {
	return &Wallet{
		db: db,
	}
}

func (w *Wallet) CreateWallet(ctx context.Context, wallet *models.Wallet) (uint64, error) {
	var id uint64
	err := w.db.QueryRowContext(
		ctx,
		"INSERT INTO public.wallets (type, coins, rubles, created_at, updated_at) VALUES ($1,$2,$3,NOW(),NOW()) RETURNING id",
		wallet.Type,
		wallet.Coins,
		wallet.Rubles,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (w *Wallet) GetWalletByID(ctx context.Context, id uint64) (*models.Wallet, error) {
	var wallet models.Wallet
	err := w.db.GetContext(ctx, &wallet, "SELECT * FROM public.wallets WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (w *Wallet) UpdateCoinsByID(ctx context.Context, id, coins uint64) error {
	_, err := w.db.ExecContext(ctx, "UPDATE public.wallets SET (coins, updated_at) = ($1, NOW()) WHERE id = $2", coins, id)
	if err != nil {
		return err
	}

	return nil
}

func (w *Wallet) UpdateRublesByID(ctx context.Context, id, rubles uint64) error {
	_, err := w.db.ExecContext(ctx, "UPDATE public.wallets SET (rubles, updated_at) = ($1, NOW()) WHERE id = $2", rubles, id)
	if err != nil {
		return err
	}

	return nil
}
