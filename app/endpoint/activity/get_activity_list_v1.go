package activity

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
)

func (a *ActivityEndpoint) GetActivityListV1(ctx context.Context, req *activity.GetActivityListV1Request) (*activity.GetActivityListV1Response, error) {
	_, err := a.services.TokenService.GetUserIDByAccessToken(ctx, req.AccessToken, a.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	list, err := a.services.ActivityService.GetActivityList(ctx)
	if err != nil {
		return nil, err
	}

	activityList := make([]*activity.Activity, 0, len(list))
	for _, value := range list {
		activityList = append(activityList, &activity.Activity{
			Id:   value.Id,
			Name: value.Name,
		})
	}

	return &activity.GetActivityListV1Response{
		ActivityList: activityList,
	}, nil
}
