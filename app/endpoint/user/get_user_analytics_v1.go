package user

import (
	"context"
	"log"

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

	analytics, err := u.services.ActivityService.GetUserActivityAnalytics(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

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
		for _, value2 := range analytics {
			if value2.Id == value.Id {
				activityList = append(activityList, &user.ActivityMessage{
					Id:   value.Id,
					Name: value.Name,
					Unit: value.Unit,
					Max:  value2.Max,
					Min:  value2.Min,
					Avg:  value2.Avg,
				})
			}
		}
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
