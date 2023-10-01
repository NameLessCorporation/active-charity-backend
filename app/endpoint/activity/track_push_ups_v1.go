package activity

import (
	"context"
	"math"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
)

func (a *ActivityEndpoint) TrackPushUpsV1(ctx context.Context, req *activity.TrackPushUpV1Request) (*activity.TrackPushUpsV1Response, error) {
	userID, err := a.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), a.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if err = a.services.ActivityService.TrackPushUps(ctx, req.Repeats, req.ActivityId, userID); err != nil {
		return nil, err
	}

	coins := uint64(math.Floor(float64(req.Repeats) / 10))

	if coins >= 1 {
		walletId, err := a.services.UserService.GetWalletIdByUserId(ctx, userID)
		if err != nil {
			return nil, err
		}

		_, err = a.operations.AccrualOfCoinsForActivity(ctx, coins, walletId)
		if err != nil {
			return nil, err
		}
	}

	return &activity.TrackPushUpsV1Response{}, nil
}
