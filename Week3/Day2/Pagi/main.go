package main

import (
	"ugdc11/config"
	"ugdc11/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()

	controller := controller.NewUserDB(db)
	g := e.Group("/user")
	{
		g.POST("/register", controller.Register)
		// g.POST("/login", controller.Login)
	}
	
	e.Start(":8080")
}