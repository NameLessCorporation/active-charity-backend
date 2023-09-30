package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) CreateOrganization(ctx context.Context, organization *models.Organization) (uint64, error) {
	id, err := s.repository.OrganizationRepository.CreateOrganization(ctx, organization)
	if err != nil {
		return 0, status.Error(codes.Internal, "Ошибка создания организации")
	}

	return id, nil
}

func (s *Service) GetOrganizationByID(ctx context.Context, id uint64) (*models.Organization, error) {
	organization, err := s.repository.OrganizationRepository.GetOrganizationByID(ctx, id)
	if err != nil {
		s.logger.Error("s.repository.OrganizationRepository.GetOrganizationByID", zap.Error(err))
		return nil, status.Error(codes.Internal, "Ошибка при получении организации")
	}

	return organization, nil
}
