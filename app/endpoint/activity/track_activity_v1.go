package activity

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
)

func (a *ActivityEndpoint) TrackActivityV1(ctx context.Context, req *activity.TrackActivityV1Request) (*activity.TrackActivityV1Response, error) {
	if req.Value == 0 {
		return &activity.TrackActivityV1Response{}, nil
	}

	switch req.ActivityId {
	case 1:
		_, err := a.TrackStepsV1(ctx, &activity.TrackStepsV1Request{
			ActivityId:  req.ActivityId,
			AccessToken: req.AccessToken,
			Steps:       req.Value,
		})
		if err != nil {
			return nil, err
		}
		break
	case 2:
		_, err := a.TrackPullUpsV1(ctx, &activity.TrackPullUpsV1Request{
			ActivityId:  req.ActivityId,
			AccessToken: req.AccessToken,
			Repeats:     req.Value,
		})
		if err != nil {
			return nil, err
		}
		break
	case 3:
		_, err := a.TrackPushUpsV1(ctx, &activity.TrackPushUpV1Request{
			ActivityId:  req.ActivityId,
			AccessToken: req.AccessToken,
			Repeats:     req.Value,
		})
		if err != nil {
			return nil, err
		}
		break
	case 4:
		_, err := a.TrackBenchPressV1(ctx, &activity.TrackBenchPressV1Request{
			ActivityId:  req.ActivityId,
			AccessToken: req.AccessToken,
			Repeats:     req.Value,
		})
		if err != nil {
			return nil, err
		}
		break
	case 5:
		_, err := a.TrackCrunchesV1(ctx, &activity.TrackCrunchesV1Request{
			ActivityId:  req.ActivityId,
			AccessToken: req.AccessToken,
			Repeats:     req.Value,
		})
		if err != nil {
			return nil, err
		}
		break
	case 6:
		_, err := a.TrackCyclingV1(ctx, &activity.TrackCyclingV1Request{
			ActivityId:  req.ActivityId,
			AccessToken: req.AccessToken,
			Metres:      req.Value,
		})
		if err != nil {
			return nil, err
		}
		break
	}
	return &activity.TrackActivityV1Response{}, nil
}
