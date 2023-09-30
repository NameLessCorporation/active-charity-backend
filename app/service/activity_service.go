package service

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (a *Service) TrackPushUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {

}

func (a *Service) TrackBenchPress(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {

}

func (a *Service) TrackCycling(ctx context.Context, metres uint32, activityId uint64, userId uint64) error {

}

func (a *Service) TrackCrunches(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {

}

func (a *Service) TrackSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) error {
	isPeriod, err := a.repository.ActivityRepository.IsActiveStepsPeriod(ctx, userId)
	if err != nil {
		return err
	}

	if isPeriod {
		periodId, err := a.repository.ActivityRepository.GetCurrentPeriodId(ctx, userId)
		if err != nil {
			return err
		}

		err = a.repository.ActivityRepository.TrackCurrentPeriodSteps(ctx, steps, periodId)
		if err != nil {
			return err
		}
	} else {
		err = a.repository.ActivityRepository.TrackNewPeriodSteps(ctx, steps, activityId, userId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Service) TrackPullUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	if err := a.repository.ActivityRepository.TrackPullUps(ctx, repeats, activityId, userId); err != nil {
		return err
	}

	return nil
}

func (a *Service) GetActivityList(ctx context.Context) ([]*models.Activity, error) {
	list, err := a.repository.ActivityRepository.GetActivityList(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
