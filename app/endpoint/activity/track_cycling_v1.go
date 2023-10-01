package activity

import (
	"context"
	"math"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
)

func (a *ActivityEndpoint) TrackCyclingV1(ctx context.Context, req *activity.TrackCyclingV1Request) (*activity.TrackCyclingV1Response, error) {
	userID, err := a.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), a.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if err = a.services.ActivityService.TrackCycling(ctx, req.Metres, req.ActivityId, userID); err != nil {
		return nil, err
	}

	coins := uint64(math.Floor(float64(req.Metres) / 10))

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

	return &activity.TrackCyclingV1Response{}, nil
}
