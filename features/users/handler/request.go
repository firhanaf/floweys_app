package handler

import (
	"floweys_app/features/users"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func RegisterRequestToCore(user RegisterRequest) users.UserCore {
	return users.UserCore{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
}
