package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) CreateTransaction(ctx context.Context, transaction *models.Transaction) (uint64, error) {
	id, err := s.repository.TransactionRepository.CreateTransaction(ctx, transaction)
	if err != nil {
		return 0, status.Error(codes.Internal, "Ошибка создания транзакции")
	}

	return id, nil
}

func (s *Service) UpdateStatusByID(ctx context.Context, id uint64, transactionStatus string) error {
	if err := s.repository.TransactionRepository.UpdateStatusByID(ctx, id, transactionStatus); err != nil {
		return status.Error(codes.Internal, "Ошибка обновления статуса через id")
	}

	return nil
}

func (s *Service) GetTransactionByID(ctx context.Context, id uint64) (*models.Transaction, error) {
	transaction, err := s.repository.TransactionRepository.GetTransactionByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения транзакции через id")
	}

	return transaction, nil
}

func (s *Service) GetTransactionByToWalletIDAndFromWalletID(ctx context.Context, fromWalletID, toWalletID uint64) (*models.Transaction, error) {
	transaction, err := s.repository.TransactionRepository.GetTransactionByToWalletIDAndFromWalletID(ctx, fromWalletID, toWalletID)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения транзакции через получателя и отправителя")
	}

	return transaction, nil
}
