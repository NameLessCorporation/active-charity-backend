package user

import (
	"context"
	"math"
	"sort"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) GetTopV1(ctx context.Context, req *user.GetTopV1Request) (*user.GetTopV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	organizationID, err := u.services.UserService.GetOrganizationByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if organizationID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Вы не состоите в организациях")
	}

	users, err := u.services.UserService.GetUsersByOrganizationID(ctx, organizationID)
	if err != nil {
		return nil, err
	}

	top := make([]*user.TopTab, 0, len(users))
	for _, v := range users {
		stepVal, err := u.services.ActivityService.GetStepsValue(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		pushVal, err := u.services.ActivityService.GetPushUpValue(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		pullValue, err := u.services.ActivityService.GetPullUpValue(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		crunVal, err := u.services.ActivityService.GetCrunchesValue(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		benchVal, err := u.services.ActivityService.GetBenchPressValue(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		cycleVal, err := u.services.ActivityService.GetCyclingValue(ctx, v.ID)
		if err != nil {
			return nil, err
		}

		points := uint64(math.Floor(float64(stepVal)/10)) +
			uint64(math.Floor(float64(cycleVal)/10)) +
			uint64(math.Floor(float64(pushVal)/10)) +
			uint64(math.Floor(float64(pullValue)/10)) +
			uint64(math.Floor(float64(crunVal)/10)) +
			uint64(math.Floor(float64(benchVal)/10))

		top = append(top, &user.TopTab{
			Name:   v.Name,
			Points: points * 100,
		})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].Points > top[j].Points
	})

	return &user.GetTopV1Response{Top: top}, nil
}
