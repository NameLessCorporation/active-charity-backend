package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) CreateWallet(ctx context.Context, wallet *models.Wallet) (uint64, error) {
	id, err := s.repository.WalletRepository.CreateWallet(ctx, wallet)
	if err != nil {
		return 0, status.Error(codes.Internal, "Ошибка создания кошелька")
	}

	return id, nil
}

func (s *Service) GetWalletByID(ctx context.Context, id uint64) (*models.Wallet, error) {
	wallet, err := s.repository.WalletRepository.GetWalletByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка при получение кошелька")
	}

	return wallet, nil
}

func (s *Service) UpdateCoinsByID(ctx context.Context, id, coins uint64) error {
	if err := s.repository.WalletRepository.UpdateCoinsByID(ctx, id, coins); err != nil {
		return status.Error(codes.Internal, "Ошибка при обновлении монет")
	}

	return nil
}

func (s *Service) UpdateRublesByID(ctx context.Context, id, rubles uint64) error {
	if err := s.repository.WalletRepository.UpdateRublesByID(ctx, id, rubles); err != nil {
		return status.Error(codes.Internal, "Ошибка при обновлении рублей")
	}

	return nil
}
