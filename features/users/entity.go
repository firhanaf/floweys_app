package users

import "time"

type UserCore struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDataInterface interface {
	Login(username, password string) (UserCore, error)
	Register(input UserCore) error
}

type UserServiceInterface interface {
	Login(username, password string) (UserCore, string, error)
	Add(input UserCore) error
}
