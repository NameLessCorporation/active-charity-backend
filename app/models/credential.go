package models

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Credential struct {
	Email    string `json:"email"    db:"email"`
	Password string `json:"password" db:"password"`
}

func (c *Credential) Validate() error {
	if len(c.Password) > 64 {
		return status.Error(codes.InvalidArgument, "Длина пароля больше 64 символов")
	}
	if len(c.Password) < 8 {
		return status.Error(codes.InvalidArgument, "Длина пароля меньше 8 символов")
	}

	return nil
}
