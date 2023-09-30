package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) CreateWallet(ctx context.Context, wallet *models.Wallet) (uint64, error) {
	id, err := s.repository.WalletRepository.CreateWallet(ctx, wallet)
	if err != nil {
		s.logger.Error("s.repository.WalletRepository.CreateWallet", zap.Error(err))
		return 0, status.Error(codes.Internal, "Ошибка создания кошелька")
	}

	return id, nil
}

func (s *Service) GetWalletByID(ctx context.Context, id uint64) (*models.Wallet, error) {
	wallet, err := s.repository.WalletRepository.GetWalletByID(ctx, id)
	if err != nil {
		s.logger.Error("s.repository.WalletRepository.GetWalletByID", zap.Error(err))
		return nil, status.Error(codes.Internal, "Ошибка при получение кошелька")
	}

	return wallet, nil
}

func (s *Service) UpdateCoinsByID(ctx context.Context, id, coins uint64) error {
	if err := s.repository.WalletRepository.UpdateCoinsByID(ctx, id, coins); err != nil {
		s.logger.Error("s.repository.WalletRepository.UpdateCoinsByID", zap.Error(err))
		return status.Error(codes.Internal, "Ошибка при обновлении монет")
	}

	return nil
}

func (s *Service) UpdateRublesByID(ctx context.Context, id, rubles uint64) error {
	if err := s.repository.WalletRepository.UpdateRublesByID(ctx, id, rubles); err != nil {
		s.logger.Error("s.repository.WalletRepository.UpdateRublesByID", zap.Error(err))
		return status.Error(codes.Internal, "Ошибка при обновлении рублей")
	}

	return nil
}
