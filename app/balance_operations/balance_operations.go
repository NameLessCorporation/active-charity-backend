package balance_operations

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/app/service"
)

type BalanceOperations interface {
	ConvertCoinsToRubles(coins uint64) uint64
	AccrualOfCoinsForActivity(ctx context.Context, coins uint64, walletID uint64) (uint64, error)
	CreateTransferCoinsToOrganization(ctx context.Context, coins uint64, fromWalletID, toWalletID uint64) (uint64, error)
	ApproveTransferCoinsToOrganization(ctx context.Context, transactionID uint64) error
	RejectTransferCoinsToOrganization(ctx context.Context, transactionID uint64) error
	WithdrawalCoinsFromOrganization(ctx context.Context, coins, walletID uint64) (uint64, error)
}

type balanceOperations struct {
	services service.Services
}

func NewBalanceOperations(services service.Services) BalanceOperations {
	return &balanceOperations{
		services: services,
	}
}

func (b *balanceOperations) ConvertCoinsToRubles(coins uint64) uint64 {
	return coins
}

func (b *balanceOperations) AccrualOfCoinsForActivity(ctx context.Context, coins uint64, walletID uint64) (uint64, error) {
	wallet, err := b.services.WalletService.GetWalletByID(ctx, walletID)
	if err != nil {
		return 0, err
	}

	var (
		rubles        = b.ConvertCoinsToRubles(coins)
		transactionID uint64
	)
	transactionID, err = b.services.TransactionService.CreateTransaction(ctx, &models.Transaction{
		Type:         models.TransactionInsertionType,
		FromWalletID: 0,
		ToWalletID:   wallet.ID,
		Coins:        coins,
		Rubles:       rubles,
		Status:       models.TransactionApprovedStatus,
	})
	if err != nil {
		return 0, err
	}

	if err = b.services.WalletService.UpdateCoinsByID(ctx, wallet.ID, wallet.Coins+coins); err != nil {
		return 0, err
	}

	if err = b.services.WalletService.UpdateRublesByID(ctx, wallet.ID, wallet.Rubles+rubles); err != nil {
		return 0, err
	}

	return transactionID, nil
}

func (b *balanceOperations) CreateTransferCoinsToOrganization(ctx context.Context, coins uint64, fromWalletID, toWalletID uint64) (uint64, error) {
	var (
		fromWallet *models.Wallet
		err        error
	)
	fromWallet, err = b.services.WalletService.GetWalletByID(ctx, fromWalletID)
	if err != nil {
		return 0, err
	}

	var rubles = b.ConvertCoinsToRubles(coins)
	if (fromWallet.Coins-coins) < 0 && (fromWallet.Rubles-rubles) < 0 {
		return 0, status.Error(codes.Internal, "Не хватает монет для перевода в организацию")
	}

	var transactionID uint64
	transactionID, err = b.services.TransactionService.CreateTransaction(ctx, &models.Transaction{
		Type:         models.TransactionTransferType,
		FromWalletID: fromWalletID,
		ToWalletID:   toWalletID,
		Coins:        coins,
		Rubles:       rubles,
		Status:       models.TransactionNewStatus,
	})
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}

func (b *balanceOperations) ApproveTransferCoinsToOrganization(ctx context.Context, transactionID uint64) error {
	var (
		transaction *models.Transaction
		err         error
	)
	transaction, err = b.services.TransactionService.GetTransactionByID(ctx, transactionID)
	if err != nil {
		return err
	}

	var fromWallet *models.Wallet
	fromWallet, err = b.services.WalletService.GetWalletByID(ctx, transaction.FromWalletID)
	if err != nil {
		return err
	}

	var toWallet *models.Wallet
	toWallet, err = b.services.WalletService.GetWalletByID(ctx, transaction.ToWalletID)
	if err != nil {
		return err
	}

	if (fromWallet.Coins-transaction.Coins) < 0 && (fromWallet.Rubles-transaction.Rubles) < 0 {
		return status.Error(codes.Internal, "Не хватает монет для перевода в организацию")
	}

	if err = b.services.WalletService.UpdateCoinsByID(ctx, transaction.FromWalletID, fromWallet.Coins-transaction.Coins); err != nil {
		return err
	}

	if err = b.services.WalletService.UpdateRublesByID(ctx, transaction.FromWalletID, fromWallet.Rubles-transaction.Rubles); err != nil {
		return err
	}

	if err = b.services.WalletService.UpdateCoinsByID(ctx, transaction.ToWalletID, toWallet.Coins+transaction.Coins); err != nil {
		return err
	}

	if err = b.services.WalletService.UpdateRublesByID(ctx, transaction.ToWalletID, toWallet.Rubles+transaction.Rubles); err != nil {
		return err
	}

	if err = b.services.TransactionService.UpdateStatusByID(ctx, transaction.ID, models.TransactionApprovedStatus); err != nil {
		return err
	}

	return nil
}

func (b *balanceOperations) RejectTransferCoinsToOrganization(ctx context.Context, transactionID uint64) error {
	transaction, err := b.services.TransactionService.GetTransactionByID(ctx, transactionID)
	if err != nil {
		return err
	}

	if err = b.services.TransactionService.UpdateStatusByID(ctx, transaction.ID, models.TransactionRejectedStatus); err != nil {
		return err
	}

	return nil
}

func (b *balanceOperations) WithdrawalCoinsFromOrganization(ctx context.Context, coins, walletID uint64) (uint64, error) {
	var (
		wallet *models.Wallet
		err    error
	)
	wallet, err = b.services.WalletService.GetWalletByID(ctx, walletID)
	if err != nil {
		return 0, err
	}

	var rubles = b.ConvertCoinsToRubles(coins)
	if (wallet.Coins-coins) < 0 && (wallet.Rubles-rubles) < 0 {
		return 0, status.Error(codes.Internal, "Не хватает монет для списания средств")
	}

	var transactionID uint64
	transactionID, err = b.services.TransactionService.CreateTransaction(ctx, &models.Transaction{
		Type:         models.TransactionWithdrawalType,
		FromWalletID: walletID,
		ToWalletID:   0,
		Coins:        coins,
		Rubles:       rubles,
		Status:       models.TransactionApprovedStatus,
	})
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}
