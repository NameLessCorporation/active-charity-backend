package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
	"github.com/NameLessCorporation/active-charity-backend/tools/helpers"
)

func (u *UserEndpoint) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserMessageResponse, error) {
	id, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	if u.services.UserService.IsExistByLogin(ctx, req.GetUserLogin()) {
		return nil, status.Error(codes.NotFound, "login is exist")
	}

	newUser := &models.User{
		Login:     req.GetUserLogin(),
		Password:  req.GetUserPassword(),
		Role:      req.GetUserRole(),
		CreatedBy: id,
	}

	if err = newUser.Validate(); err != nil {
		return nil, err
	}

	hashPassword, err := helpers.NewMD5Hash(req.GetUserPassword())
	if err != nil {
		return nil, err
	}

	newUser.Password = hashPassword

	_, err = u.services.UserService.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &user.UserMessageResponse{
		Message: "user was created successfully",
	}, nil
}
