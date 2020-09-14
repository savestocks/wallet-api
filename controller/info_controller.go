package controller

import (
	"net/http"

	"github.com/andersonlira/wallet-api/config"
	"github.com/andersonlira/wallet-api/domain"
	"github.com/labstack/echo/v4"
)

//GetInfo of the application like version
func GetInfo(c echo.Context) error {
	info := domain.Info{Version: config.Version}
	return c.JSON(http.StatusOK, info)
}
