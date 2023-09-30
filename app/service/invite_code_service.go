package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) CreateInviteCode(ctx context.Context, inviteCode *models.InviteCode) (uint64, error) {
	id, err := s.repository.InviteCodeRepository.CreateInviteCode(ctx, inviteCode)
	if err != nil {
		s.logger.Error("s.repository.InviteCodeRepository.CreateInviteCode", zap.Error(err))
		return 0, status.Error(codes.Internal, "Ошибка создания организации")
	}

	return id, nil
}

func (s *Service) GetInviteCodeByCode(ctx context.Context, code string) (*models.InviteCode, error) {
	inviteCode, err := s.repository.InviteCodeRepository.GetInviteCodeByCode(ctx, code)
	if err != nil {
		s.logger.Error("s.repository.InviteCodeRepository.GetInviteCodeByCode", zap.Error(err))
		return nil, status.Error(codes.Internal, "Ошибка получения кода приглашения от организации")
	}

	return inviteCode, nil
}
