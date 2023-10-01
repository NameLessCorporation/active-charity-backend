package user

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) GetUserAnalyticsV1(ctx context.Context, req *user.GetUserAnalyticsV1Request) (*user.GetUserAnalyticsV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	list, err := u.services.ActivityService.GetActivityList(ctx)
	if err != nil {
		return nil, err
	}

	u.services.ActivityService.GetUserActivityAnalytics(ctx, userID, list)

	fav, err := u.services.ActivityService.GetUserFavouriteActivity(ctx, userID)
	if err != nil {
		return nil, err
	}

	mostEarned, err := u.services.ActivityService.GetUserMostEarnedActivity(ctx, userID)
	if err != nil {
		return nil, err
	}

	activityList := make([]*user.ActivityMessage, 0, len(list))
	for _, value := range list {
		activityList = append(activityList, &user.ActivityMessage{
			Id:   value.Id,
			Name: value.Name,
			Unit: value.Unit,
			Max:  1,
			Min:  1,
			Avg:  1,
		})
	}

	return &user.GetUserAnalyticsV1Response{
		ActivityList:       activityList,
		MostEarnedActivity: mostEarned,
		FavouriteActivity:  fav,
		Steps:              nil,
		PushUps:            nil,
		Crunches:           nil,
		Cycling:            nil,
		PullUps:            nil,
		BenchPress:         nil,
	}, nil

}
