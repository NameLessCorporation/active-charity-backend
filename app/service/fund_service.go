package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) CreateFund(ctx context.Context, fund *models.Fund) (uint64, error) {
	id, err := s.repository.FundRepository.CreateFund(ctx, fund)
	if err != nil {
		return 0, status.Error(codes.Internal, "Ошибка создания фонда")
	}

	return id, nil
}

func (s *Service) GetFundByID(ctx context.Context, id uint64) (*models.Fund, error) {
	fund, err := s.Services.FundService.GetFundByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения фонда")
	}

	return fund, nil
}

func (s *Service) GetFunds(ctx context.Context) ([]*models.Fund, error) {
	funds, err := s.Services.FundService.GetFunds(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения фондов")
	}

	return funds, nil
}
