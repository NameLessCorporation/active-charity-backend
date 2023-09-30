package endpoint

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
	"github.com/NameLessCorporation/active-charity-backend/extra/auth"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

type EndpointContainer struct {
	AuthService         AuthServiceInter
	UserService         UserServiceInter
	OrganizationService OrganizationServiceInter
	ActivityService     ActivityServiceInter
}

func NewEndpointContainer(
	auth AuthServiceInter,
	user UserServiceInter,
	organizationService OrganizationServiceInter,
	activity ActivityServiceInter,
) *EndpointContainer {
	return &EndpointContainer{
		AuthService:         auth,
		UserService:         user,
		OrganizationService: organizationService,
		ActivityService:     activity,
	}
}

type AuthServiceInter interface {
	AuthWithCredentialsV1(ctx context.Context, req *auth.AuthWithCredentialsV1Request) (*auth.AccessRefreshTokens, error)
	RefreshTokensV1(ctx context.Context, req *auth.AccessRefreshTokens) (*auth.AccessRefreshTokens, error)
	AuthWithAccessTokenV1(ctx context.Context, req *auth.AuthWithAccessTokenV1Request) (*auth.AuthWithAccessTokenV1Response, error)
}

type UserServiceInter interface {
	CreateUserV1(ctx context.Context, req *user.CreateUserV1Request) (*user.CreateUserV1Response, error)
	JoinUserToOrganizationV1(ctx context.Context, req *user.JoinUserToOrganizationV1Request) (*user.JoinUserToOrganizationV1Response, error)
}

type OrganizationServiceInter interface {
	CreateOrganizationV1(ctx context.Context, req *organization.CreateOrganizationV1Request) (*organization.CreateOrganizationV1Response, error)
	CreateOrganizationInviteCodeV1(ctx context.Context, req *organization.CreateOrganizationInviteCodeV1Request) (*organization.CreateOrganizationInviteCodeV1Response, error)
}

type ActivityServiceInter interface {
	TrackStepsV1(ctx context.Context, req *activity.TrackStepsV1Request) (*activity.TrackStepsV1Response, error)
	TrackBenchPressV1(ctx context.Context, req *activity.TrackBenchPressV1Request) (*activity.TrackBenchPressV1Response, error)
	TrackPullUpsV1(ctx context.Context, req *activity.TrackPullUpsV1Request) (*activity.TrackPullUpsV1Response, error)
	TrackCyclingV1(ctx context.Context, req *activity.TrackCyclingV1Request) (*activity.TrackCyclingV1Response, error)
	TrackCrunchesV1(ctx context.Context, req *activity.TrackCrunchesV1Request) (*activity.TrackCrunchesV1Response, error)
	TrackPushUpsV1(ctx context.Context, req *activity.TrackPushUpV1Request) (*activity.TrackPushUpsV1Response, error)

	GetActivityListV1(ctx context.Context, req *activity.GetActivityListV1Request) (*activity.GetActivityListV1Response, error)
}
