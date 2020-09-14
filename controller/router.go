package controller

import (
	"github.com/andersonlira/wallet-api/config"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//MapRoutes for the endpoints which the API listens for
func MapRoutes(e *echo.Echo) {
	g := e.Group("/wallet-api/v1")
	if config.Values.UsePrometheus {
		p := prometheus.NewPrometheus("echo", nil)
		p.Use(e)
	}
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	g.GET("/expense", GetExpenseList)
	g.GET("/expense/:id", GetExpenseByID)
	g.POST("/expense", SaveExpense)
	g.PUT("/expense/:id", UpdateExpense)
	g.DELETE("/expense/:id", DeleteExpense)
	g.GET("/walletPosition", GetWalletPositionList)
	g.GET("/walletPosition/:id", GetWalletPositionByID)
	g.POST("/walletPosition", SaveWalletPosition)
	g.PUT("/walletPosition/:id", UpdateWalletPosition)
	g.DELETE("/walletPosition/:id", DeleteWalletPosition)
	g.GET("/health", CheckHealth)
	g.GET("/info", GetInfo)
}


