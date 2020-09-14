package controller

import (
    "net/http"

    "github.com/andersonlira/wallet-api/gateway/txtdb"
    "github.com/andersonlira/wallet-api/domain"
	"github.com/labstack/echo/v4"

)


//GetWalletPositionList return all objects 
func GetWalletPositionList(c echo.Context) error {

    list := txtdb.GetWalletPositionList()

	return c.JSON(http.StatusOK, list)
}

func GetWalletPositionByID(c echo.Context) error {
    ID := c.Param("id")
    it, err := txtdb.GetWalletPositionByID(ID)
    if err != nil {
        return c.JSON(http.StatusNotFound,it)
    }
    return c.JSON(http.StatusOK, it)
}

func SaveWalletPosition(c echo.Context) error {
    it := domain.WalletPosition{}
    c.Bind(&it)
    it = txtdb.SaveWalletPosition(it)
    return c.JSON(http.StatusCreated, it)
}

func UpdateWalletPosition(c echo.Context) error {
    ID := c.Param("id")
    it := domain.WalletPosition{}
    c.Bind(&it)
    it = txtdb.UpdateWalletPosition(ID,it)
    return c.JSON(http.StatusOK, it)
}

func DeleteWalletPosition(c echo.Context) error {
    ID := c.Param("id")
    result := txtdb.DeleteWalletPosition(ID)
    return c.JSON(http.StatusOK, result)
}