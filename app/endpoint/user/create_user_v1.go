package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
	"github.com/NameLessCorporation/active-charity-backend/tools/helpers"
)

func (u *UserEndpoint) CreateUserV1(ctx context.Context, req *user.CreateUserV1Request) (*user.CreateUserV1Response, error) {
	if u.services.UserService.IsExistByEmail(ctx, req.GetEmail()) {
		return nil, status.Error(codes.AlreadyExists, "Данный email уже зарегистрирован")
	}

	newUser := &models.User{
		Email:          req.GetEmail(),
		Password:       req.GetPassword(),
		Name:           req.GetName(),
		DateOfBirthday: req.GetDateOfBirthday(),
	}
	if err := newUser.Validate(); err != nil {
		return nil, err
	}

	hashPassword, err := helpers.NewMD5Hash(req.GetPassword())
	if err != nil {
		return nil, err
	}

	newUser.Password = hashPassword

	_, err = u.services.UserService.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserV1Response{}, nil
}
