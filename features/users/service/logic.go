package service

import (
	"errors"
	"floweys_app/app/middlewares"
	"floweys_app/features/users"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

func (service *UserService) Add(input users.UserCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate, required data")
	}
	if len(input.Password) < 8 {
		return errors.New("password length min. 8 characters")
	}
	errInsert := service.userData.Register(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (service *UserService) Login(username, password string) (dataLogin users.UserCore, token string, err error) {
	dataLogin, err = service.userData.Login(username, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil

}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
		validate: validator.New(),
	}
}
