package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (s *Service) TrackSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) error {
	err := s.repository.ActivityRepository.TrackSteps(ctx, steps, activityId, userId)
	if err != nil {
		s.logger.Error("s.repository.ActivityRepository.TrackSteps", zap.Error(err))
		return err
	}

	// req to earn coin

	return nil
}

func (s *Service) GetActivityList(ctx context.Context) ([]*models.Activity, error) {
	list, err := s.repository.ActivityRepository.GetActivityList(ctx)
	if err != nil {
		s.logger.Error("s.repository.ActivityRepository.GetActivityList", zap.Error(err))
		return nil, err
	}

	return list, nil
}
