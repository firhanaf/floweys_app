package router

import (
	"floweys_app/app/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_userData "floweys_app/features/users/data"
	_userHandler "floweys_app/features/users/handler"
	_userService "floweys_app/features/users/service"
)

func InitRouter(db *gorm.DB, c *echo.Echo, cfg *config.AppConfig) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	c.POST("/login", UserHandlerAPI.Login)
	c.POST("/register", UserHandlerAPI.Register)

}
