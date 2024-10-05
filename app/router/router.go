package router

import (
	"floweys_app/app/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_userData "floweys_app/features/users/data"
	_userHandler "floweys_app/features/users/handler"
	_userService "floweys_app/features/users/service"

	_orderData "floweys_app/features/orders/data"
	_orderHandler "floweys_app/features/orders/handler"
	_orderService "floweys_app/features/orders/service"
)

func InitRouter(db *gorm.DB, c *echo.Echo, cfg *config.AppConfig) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	OrderData := _orderData.New(db)
	OrderService := _orderService.New(OrderData)
	OrderHandlerAPI := _orderHandler.New(OrderService)

	c.POST("/login", UserHandlerAPI.Login)
	c.POST("/register", UserHandlerAPI.Register)
	c.POST("/order", OrderHandlerAPI.Create)

}
