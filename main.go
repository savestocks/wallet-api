package main

import (
	"github.com/andersonlira/wallet-api/config"
	"github.com/andersonlira/wallet-api/controller"
	_ "github.com/andersonlira/wallet-api/gateway/customlog"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
