package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) CreateTransaction(ctx context.Context, transaction *models.Transaction) (uint64, error) {
	id, err := s.repository.TransactionRepository.CreateTransaction(ctx, transaction)
	if err != nil {
		s.logger.Error("s.repository.TransactionRepository.CreateTransaction", zap.Error(err))
		return 0, status.Error(codes.Internal, "Ошибка создания транзакции")
	}

	return id, nil
}

func (s *Service) UpdateStatusByID(ctx context.Context, id uint64, transactionStatus string) error {
	if err := s.repository.TransactionRepository.UpdateStatusByID(ctx, id, transactionStatus); err != nil {
		s.logger.Error("s.repository.TransactionRepository.UpdateStatusByID", zap.Error(err))
		return status.Error(codes.Internal, "Ошибка обновления статуса через id")
	}

	return nil
}

func (s *Service) GetTransactionByID(ctx context.Context, id uint64) (*models.Transaction, error) {
	transaction, err := s.repository.TransactionRepository.GetTransactionByID(ctx, id)
	if err != nil {
		s.logger.Error("s.repository.TransactionRepository.GetTransactionByID", zap.Error(err))
		return nil, status.Error(codes.Internal, "Ошибка получения транзакции через id")
	}

	return transaction, nil
}

func (s *Service) GetTransactionByToWalletIDAndFromWalletID(ctx context.Context, fromWalletID, toWalletID uint64) (*models.Transaction, error) {
	transaction, err := s.repository.TransactionRepository.GetTransactionByToWalletIDAndFromWalletID(ctx, fromWalletID, toWalletID)
	if err != nil {
		s.logger.Error("s.repository.TransactionRepository.GetTransactionByToWalletIDAndFromWalletID", zap.Error(err))
		return nil, status.Error(codes.Internal, "Ошибка получения транзакции через получателя и отправителя")
	}

	return transaction, nil
}

func (s *Service) GetNewTransferTransactionsByToWalletID(ctx context.Context, toWalletID uint64) ([]*models.Transaction, error) {
	transactions, err := s.repository.TransactionRepository.GetNewTransferTransactionsByToWalletID(ctx, toWalletID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения новых трансферов")
	}

	return transactions, nil
}

func (s *Service) GetNewTransferTransactionsByWalletID(ctx context.Context, walletID uint64, transactionType, transactionStatus string) ([]*models.Transaction, error) {
	transactions, err := s.repository.TransactionRepository.GetTransactionsByWalletID(ctx, walletID, transactionType, transactionStatus)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения транзакций")
	}

	return transactions, nil
}
