package service

import (
	"context"
	"log"

	"go.uber.org/zap"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (a *Service) TrackPushUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	if err := a.repository.ActivityRepository.TrackPushUps(ctx, repeats, activityId, userId); err != nil {
		return err
	}

	return nil
}

func (a *Service) TrackBenchPress(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	if err := a.repository.ActivityRepository.TrackBenchPress(ctx, repeats, activityId, userId); err != nil {
		return err
	}

	return nil
}

func (a *Service) TrackCycling(ctx context.Context, metres uint32, activityId uint64, userId uint64) error {
	if err := a.repository.ActivityRepository.TrackCycling(ctx, metres, activityId, userId); err != nil {
		return err
	}

	return nil
}

func (a *Service) TrackCrunches(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	if err := a.repository.ActivityRepository.TrackCrunches(ctx, repeats, activityId, userId); err != nil {
		return err
	}

	return nil
}

func (a *Service) TrackSteps(ctx context.Context, steps uint32, activityId uint64, userId uint64) (uint32, error) {
	isPeriod, err := a.repository.ActivityRepository.IsActiveStepsPeriod(ctx, userId)
	if err != nil {
		a.logger.Error("s.repository.ActivityRepository.TrackSteps", zap.Error(err))
		return 0, err
	}

	if isPeriod {
		periodId, value, err := a.repository.ActivityRepository.GetCurrentPeriodId(ctx, userId)
		if err != nil {
			return 0, err
		}

		err = a.repository.ActivityRepository.TrackCurrentPeriodSteps(ctx, steps, periodId)
		if err != nil {
			return 0, err
		}

		log.Println("PIZDEC", steps-value)

		return steps - value, nil
	} else {
		err = a.repository.ActivityRepository.TrackNewPeriodSteps(ctx, steps, activityId, userId)
		if err != nil {
			return 0, err
		}

		return steps, nil
	}
}

func (a *Service) TrackPullUps(ctx context.Context, repeats uint32, activityId uint64, userId uint64) error {
	if err := a.repository.ActivityRepository.TrackPullUps(ctx, repeats, activityId, userId); err != nil {
		return err
	}

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
