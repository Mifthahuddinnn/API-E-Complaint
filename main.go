package main

import (
	"api-e-complaint/config"
	"api-e-complaint/drivers/mysql"
	"api-e-complaint/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.LoadEnv()
	config.InitConfigMySQL()
	DB := mysql.ConnectDB(config.InitConfigMySQL())
	e := echo.New()

	fmt.Println(DB)

	routes := routes.RouteController{}

	routes.InitRoute(e)
	e.Start(":8080")
}
