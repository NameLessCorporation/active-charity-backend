package activity

import (
	"context"
	"math"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
)

func (a *ActivityEndpoint) TrackStepsV1(ctx context.Context, req *activity.TrackStepsV1Request) (*activity.TrackStepsV1Response, error) {
	userID, err := a.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), a.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if err = a.services.ActivityService.TrackSteps(ctx, req.Steps, req.ActivityId, userID); err != nil {
		return nil, err
	}

	walletId, err := a.services.UserService.GetWalletIdByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	coins := uint64(math.Floor(float64(req.Steps) / 10))

	_, err = a.operations.AccrualOfCoinsForActivity(ctx, coins, walletId)
	if err != nil {
		return nil, err
	}

	return &activity.TrackStepsV1Response{}, nil
}
