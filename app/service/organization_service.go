package service

import (
	"context"

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
