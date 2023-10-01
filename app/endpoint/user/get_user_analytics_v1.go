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

	analytics, err := u.services.ActivityService.GetUserActivityAnalytics(ctx, userID)
	if err != nil {
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

	pullGraph, err := u.services.ActivityService.GetPullUpGraph(ctx, userID)
	if err != nil {
		return nil, err
	}
	pullGraphResp := make([]*user.GraphElement, 0, len(pullGraph))
	for _, graph := range pullGraph {
		pullGraphResp = append(pullGraphResp, &user.GraphElement{
			Timestamp: graph.Timestamp.UnixNano(),
			Value:     graph.Value,
		})
	}

	cycleGraph, err := u.services.ActivityService.GetCycleGraph(ctx, userID)
	if err != nil {
		return nil, err
	}
	cycleGraphResp := make([]*user.GraphElement, 0, len(cycleGraph))
	for _, graph := range cycleGraph {
		cycleGraphResp = append(cycleGraphResp, &user.GraphElement{
			Timestamp: graph.Timestamp.UnixNano(),
			Value:     graph.Value,
		})
	}

	crunchGraph, err := u.services.ActivityService.GetCrunchesGraph(ctx, userID)
	if err != nil {
		return nil, err
	}
	crunchGraphResp := make([]*user.GraphElement, 0, len(crunchGraph))
	for _, graph := range crunchGraph {
		crunchGraphResp = append(crunchGraphResp, &user.GraphElement{
			Timestamp: graph.Timestamp.UnixNano(),
			Value:     graph.Value,
		})
	}

	stepsGraph, err := u.services.ActivityService.GetStepsGraph(ctx, userID)
	if err != nil {
		return nil, err
	}
	stepsGraphResp := make([]*user.GraphElement, 0, len(stepsGraph))
	for _, graph := range stepsGraph {
		stepsGraphResp = append(stepsGraphResp, &user.GraphElement{
			Timestamp: graph.Timestamp.UnixNano(),
			Value:     graph.Value,
		})
	}

	pushGraph, err := u.services.ActivityService.GetPushUpGraph(ctx, userID)
	if err != nil {
		return nil, err
	}
	pushGraphResp := make([]*user.GraphElement, 0, len(pushGraph))
	for _, graph := range pushGraph {
		pushGraphResp = append(pushGraphResp, &user.GraphElement{
			Timestamp: graph.Timestamp.UnixNano(),
			Value:     graph.Value,
		})
	}

	benchGraph, err := u.services.ActivityService.GetBenchPressGraph(ctx, userID)
	if err != nil {
		return nil, err
	}
	benchGraphResp := make([]*user.GraphElement, 0, len(benchGraph))
	for _, graph := range benchGraph {
		benchGraphResp = append(benchGraphResp, &user.GraphElement{
			Timestamp: graph.Timestamp.UnixNano(),
			Value:     graph.Value,
		})
	}

	return &user.GetUserAnalyticsV1Response{
		ActivityList:       activityList,
		MostEarnedActivity: mostEarned,
		FavouriteActivity:  fav,
		Steps:              stepsGraphResp,
		PushUps:            pushGraphResp,
		Crunches:           crunchGraphResp,
		Cycling:            cycleGraphResp,
		PullUps:            pushGraphResp,
		BenchPress:         benchGraphResp,
	}, nil

}
