package activity

import (
	"context"
	"math"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
)

func (a *ActivityEndpoint) TrackBenchPressV1(ctx context.Context, req *activity.TrackBenchPressV1Request) (*activity.TrackBenchPressV1Response, error) {
	userID, err := a.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), a.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if err = a.services.ActivityService.TrackBenchPress(ctx, req.Repeats, req.ActivityId, userID); err != nil {
		return nil, err
	}

	walletId, err := a.services.UserService.GetWalletIdByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	coins := uint64(math.Floor(float64(req.Repeats) / 10))

	_, err = a.operations.AccrualOfCoinsForActivity(ctx, coins, walletId)
	if err != nil {
		return nil, err
	}

	return &activity.TrackBenchPressV1Response{}, nil
}
