package database

import (
	"floweys_app/app/config"
	usersData "floweys_app/features/users/data"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySql(config *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUsername, config.DBPass, config.DBHost, config.DBPORT, config.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&usersData.User{})
}
