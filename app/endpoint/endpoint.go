package endpoint

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/extra/activity"
	"github.com/NameLessCorporation/active-charity-backend/extra/auth"
	"github.com/NameLessCorporation/active-charity-backend/extra/fund"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

type EndpointContainer struct {
	AuthService         AuthServiceInter
	UserService         UserServiceInter
	OrganizationService OrganizationServiceInter
	ActivityService     ActivityServiceInter
	FundService         FundServiceInter
}

func NewEndpointContainer(
	auth AuthServiceInter,
	user UserServiceInter,
	organizationService OrganizationServiceInter,
	activity ActivityServiceInter,
	fund FundServiceInter,
) *EndpointContainer {
	return &EndpointContainer{
		AuthService:         auth,
		UserService:         user,
		OrganizationService: organizationService,
		ActivityService:     activity,
		FundService:         fund,
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
	GetUserV1(ctx context.Context, req *user.GetUserV1Request) (*user.GetUserV1Response, error)
	SelectUserFundV1(ctx context.Context, req *user.SelectUserFundV1Request) (*user.SelectUserFundV1Response, error)
	CreateNewTransferV1(ctx context.Context, req *user.CreateNewTransferV1Request) (*user.CreateNewTransferV1Response, error)
	GetTransactionsV1(ctx context.Context, req *user.GetTransactionsV1Request) (*user.GetTransactionsV1Response, error)
	GetTopV1(ctx context.Context, req *user.GetTopV1Request) (*user.GetTopV1Response, error)
	GetUserAnalyticsV1(ctx context.Context, req *user.GetUserAnalyticsV1Request) (*user.GetUserAnalyticsV1Response, error)
}

type OrganizationServiceInter interface {
	CreateOrganizationV1(ctx context.Context, req *organization.CreateOrganizationV1Request) (*organization.CreateOrganizationV1Response, error)
	CreateOrganizationInviteCodeV1(ctx context.Context, req *organization.CreateOrganizationInviteCodeV1Request) (*organization.CreateOrganizationInviteCodeV1Response, error)
	GetOrganizationV1(ctx context.Context, req *organization.GetOrganizationV1Request) (*organization.GetOrganizationV1Response, error)
	ApproveTransferCoinsV1(ctx context.Context, req *organization.ApproveTransferCoinsV1Request) (*organization.ApproveTransferCoinsV1Response, error)
	RejectTransferCoinsV1(ctx context.Context, req *organization.RejectTransferCoinsV1Request) (*organization.RejectTransferCoinsV1Response, error)
	GetNewTransfersV1(ctx context.Context, req *organization.GetNewTransfersV1Request) (*organization.GetNewTransfersV1Response, error)
	WithdrawalCoinsV1(ctx context.Context, req *organization.WithdrawalCoinsV1Request) (*organization.WithdrawalCoinsV1Response, error)
	GetTransactionsV1(ctx context.Context, req *organization.GetTransactionsV1Request) (*organization.GetTransactionsV1Response, error)
	GetOrganizationUsersV1(ctx context.Context, req *organization.GetOrganizationUsersV1Request) (*organization.GetOrganizationUsersV1Response, error)
	GetOrganizationUserAnalyticsV1(ctx context.Context, req *organization.GetOrganizationUserAnalyticsV1Request) (*organization.GetOrganizationUserAnalyticsV1Response, error)
}

type ActivityServiceInter interface {
	TrackStepsV1(ctx context.Context, req *activity.TrackStepsV1Request) (*activity.TrackStepsV1Response, error)
	TrackBenchPressV1(ctx context.Context, req *activity.TrackBenchPressV1Request) (*activity.TrackBenchPressV1Response, error)
	TrackPullUpsV1(ctx context.Context, req *activity.TrackPullUpsV1Request) (*activity.TrackPullUpsV1Response, error)
	TrackCyclingV1(ctx context.Context, req *activity.TrackCyclingV1Request) (*activity.TrackCyclingV1Response, error)
	TrackCrunchesV1(ctx context.Context, req *activity.TrackCrunchesV1Request) (*activity.TrackCrunchesV1Response, error)
	TrackPushUpsV1(ctx context.Context, req *activity.TrackPushUpV1Request) (*activity.TrackPushUpsV1Response, error)
	TrackActivityV1(ctx context.Context, req *activity.TrackActivityV1Request) (*activity.TrackActivityV1Response, error)
	GetActivityListV1(ctx context.Context, req *activity.GetActivityListV1Request) (*activity.GetActivityListV1Response, error)
}

type FundServiceInter interface {
	CreateFundV1(ctx context.Context, req *fund.CreateFundV1Request) (*fund.CreateFundV1Response, error)
	GetFundV1(ctx context.Context, req *fund.GetFundV1Request) (*fund.GetFundV1Response, error)
	GetFundsV1(ctx context.Context, req *fund.GetFundsV1Request) (*fund.GetFundsV1Response, error)
}
