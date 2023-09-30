package service

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (a *Service) TrackSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) error {
	err := a.repository.ActivityRepository.TrackSteps(ctx, steps, activityId, userId)
	if err != nil {
		return err
	}

	// req to earn coin

	return nil
}

func (a *Service) GetActivityList(ctx context.Context) ([]*models.Activity, error) {
	list, err := a.repository.ActivityRepository.GetActivityList(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
