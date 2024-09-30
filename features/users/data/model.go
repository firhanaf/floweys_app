package data

import (
	"floweys_app/features/users"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username;unique;not null"`
	Password string `gorm:"password;not null"`
	Email    string `gorm:"email;unique;not null"`
	Phone    string `gorm:"phone;unique;not null"`
}

func UserCoreToModel(input users.UserCore) User {
	return User{
		Model:    gorm.Model{},
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}

func UserModelToCore(input User) users.UserCore {
	return users.UserCore{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}
