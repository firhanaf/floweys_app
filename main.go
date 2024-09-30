package main

import (
	"floweys_app/app/config"
	"floweys_app/app/database"
	"floweys_app/app/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbMySql := database.InitMySql(cfg)
	database.InitialMigration(dbMySql)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(dbMySql, e, cfg)
	e.Logger.Fatal(e.Start(":80"))

}
