package data

import (
	"errors"
	"floweys_app/app/helpers"
	"floweys_app/features/users"
	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

func (repo *UserQuery) Register(input users.UserCore) error {
	var userModel = UserCoreToModel(input)

	hash, errHash := helpers.HashPassword(userModel.Password)
	if errHash != nil {
		return errHash
	}
	userModel.Password = hash

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *UserQuery) Login(username string, password string) (dataLogin users.UserCore, err error) {
	var data User
	tx := repo.db.Where("username = ?", username).First(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}

	passwordCheck := helpers.CheckPassword(password, data.Password)
	if !passwordCheck {
		return users.UserCore{}, helpers.ErrLogin
	}

	if tx.RowsAffected == 0 {
		return users.UserCore{}, err
	}
	dataLogin = UserModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db:        db,
		dataLogin: users.UserCore{},
	}
}
