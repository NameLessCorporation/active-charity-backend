package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) IsExistByEmail(ctx context.Context, email string) bool {
	return s.repository.UserRepository.IsExistByEmail(ctx, email)
}

func (s *Service) GetWalletIdByUserId(ctx context.Context, userId uint64) (uint64, error) {
	walletId, err := s.repository.UserRepository.GetWalletIdById(ctx, userId)
	if err != nil {
		return 0, err
	}

	return walletId, nil
}

func (s *Service) GetIDByCredentials(ctx context.Context, credentials *models.Credential) (uint64, error) {
	id, err := s.repository.UserRepository.GetIDByCredentials(ctx, credentials)
	if err != nil {
		return 0, status.Error(codes.Internal, "Ошибка получения id пользователя через учетные данные")
	}

	return id, nil
}

func (s *Service) GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	user, err := s.repository.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Ошибка получения пользователя через id")
	}

	return user, nil
}

func (s *Service) CreateUser(ctx context.Context, user *models.User) (uint64, error) {
	id, err := s.repository.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return 0, status.Error(codes.Internal, "Ошибка создания пользователя")
	}

	return id, nil
}

func (s *Service) GetIDByEmail(ctx context.Context, email string) (uint64, error) {
	id, err := s.repository.UserRepository.GetIDByEmail(ctx, email)
	if err != nil {
		return 0, status.Error(codes.Internal, "Ошибка получения email пользователя через id")
	}

	return id, nil
}

func (s *Service) UpdateOrganizationIDByID(ctx context.Context, id, organizationID uint64) error {
	if err := s.repository.UserRepository.UpdateOrganizationIDByID(ctx, id, organizationID); err != nil {
		return status.Error(codes.Internal, "Ошибка привязки к организации")
	}

	return nil
}
