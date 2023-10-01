package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
)

func (a *Service) GetStepsGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	graph, err := a.repository.ActivityRepository.GetStepsGraph(ctx, userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Service) GetPushUpGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	graph, err := a.repository.ActivityRepository.GetPushUpGraph(ctx, userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Service) GetCrunchesGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	graph, err := a.repository.ActivityRepository.GetCrunchesGraph(ctx, userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Service) GetCycleGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	graph, err := a.repository.ActivityRepository.GetCycleGraph(ctx, userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Service) GetPullUpGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	graph, err := a.repository.ActivityRepository.GetPullUpGraph(ctx, userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Service) GetBenchPressGraph(ctx context.Context, userID uint64) ([]*models.Graph, error) {
	graph, err := a.repository.ActivityRepository.GetBenchPressGraph(ctx, userID)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (a *Service) GetUserActivityAnalytics(ctx context.Context, userID uint64) ([]models.ActivityAnalytics, error) {
	l, err := a.repository.ActivityRepository.GetUserActivityAnalytics(ctx, userID)
	if err != nil {
		return nil, err
	}

	return l, err
}

func (a *Service) GetUserFavouriteActivity(ctx context.Context, userID uint64) (string, error) {
	activity, err := a.repository.ActivityRepository.GetUserFavouriteActivity(ctx, userID)
	if err != nil {
		return "", err
	}

	return activity, nil
}

func (a *Service) GetUserMostEarnedActivity(ctx context.Context, userID uint64) (string, error) {
	activity, err := a.repository.ActivityRepository.GetUserMostEarnedActivity(ctx, userID)
	if err != nil {
		return "", err
	}

	return activity, nil
}

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

func (s *Service) GetStepsValue(ctx context.Context, userID uint64) (uint32, error) {
	value, err := s.repository.ActivityRepository.GetStepsValue(ctx, userID)
	if err != nil {
		return 0, err
	}

	return uint32(value.Int32), nil
}

func (s *Service) GetBenchPressValue(ctx context.Context, userID uint64) (uint32, error) {
	value, err := s.repository.ActivityRepository.GetBenchPressValue(ctx, userID)
	if err != nil {
		return 0, err
	}

	return uint32(value.Int32), nil
}

func (s *Service) GetCrunchesValue(ctx context.Context, userID uint64) (uint32, error) {
	value, err := s.repository.ActivityRepository.GetCrunchesValue(ctx, userID)
	if err != nil {
		return 0, err
	}

	return uint32(value.Int32), nil
}

func (s *Service) GetCyclingValue(ctx context.Context, userID uint64) (uint32, error) {
	value, err := s.repository.ActivityRepository.GetCyclingValue(ctx, userID)
	if err != nil {
		return 0, err
	}

	return uint32(value.Int32), nil
}

func (s *Service) GetPullUpValue(ctx context.Context, userID uint64) (uint32, error) {
	value, err := s.repository.ActivityRepository.GetPullUpValue(ctx, userID)
	if err != nil {
		return 0, err
	}

	return uint32(value.Int32), nil
}

func (s *Service) GetPushUpValue(ctx context.Context, userID uint64) (uint32, error) {
	value, err := s.repository.ActivityRepository.GetPushUpValue(ctx, userID)
	if err != nil {
		return 0, err
	}

	return uint32(value.Int32), nil
}
