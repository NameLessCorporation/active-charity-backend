package activity

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
)

func (a *ActivityEndpoint) TrackStepsV1(ctx context.Context, req *activity.TrackStepsV1Request) (*activity.TrackStepsV1Response, error) {
	userID, err := a.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), a.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if err = a.services.ActivityService.TrackSteps(ctx, req.StepsPerDay, req.ActivityId, userID); err != nil {
		return nil, err
	}

	return &activity.TrackStepsV1Response{}, nil
}
