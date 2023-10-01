package organization

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
)

func (o *OrganizationEndpoint) GetOrganizationUserAnalyticsV1(ctx context.Context, req *organization.GetOrganizationUserAnalyticsV1Request) (*organization.GetOrganizationUserAnalyticsV1Response, error) {
	userID, err := o.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), o.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var usr *models.User
	usr, err = o.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var org *models.Organization
	org, err = o.services.OrganizationService.GetOrganizationByID(ctx, usr.OrganizationID)

	if org.OwnerID != userID {
		return nil, status.Error(codes.PermissionDenied, "Пользователь не владеет организацией")
	}

	list, err := o.services.ActivityService.GetActivityList(ctx)
	if err != nil {
		return nil, err
	}

	analytics, err := o.services.ActivityService.GetUserActivityAnalytics(ctx, req.GetUserId())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fav, err := o.services.ActivityService.GetUserFavouriteActivity(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	mostEarned, err := o.services.ActivityService.GetUserMostEarnedActivity(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	activityList := make([]*organization.ActivityMessage, 0, len(list))
	for _, value := range list {
		for _, value2 := range analytics {
			if value2.Id == value.Id {
				activityList = append(activityList, &organization.ActivityMessage{
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

	return &organization.GetOrganizationUserAnalyticsV1Response{
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
