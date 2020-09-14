package controller

import (
    "net/http"

    "github.com/andersonlira/wallet-api/gateway/txtdb"
    "github.com/andersonlira/wallet-api/domain"
	"github.com/labstack/echo/v4"

)


//GetBudgetList return all objects 
func GetBudgetList(c echo.Context) error {

    list := txtdb.GetBudgetList()

	return c.JSON(http.StatusOK, list)
}

func GetBudgetByID(c echo.Context) error {
    ID := c.Param("id")
    it, err := txtdb.GetBudgetByID(ID)
    if err != nil {
        return c.JSON(http.StatusNotFound,it)
    }
    return c.JSON(http.StatusOK, it)
}

func SaveBudget(c echo.Context) error {
    it := domain.Budget{}
    c.Bind(&it)
    it = txtdb.SaveBudget(it)
    return c.JSON(http.StatusCreated, it)
}

func UpdateBudget(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Budget{}
    c.Bind(&it)
    it = txtdb.UpdateBudget(ID,it)
    return c.JSON(http.StatusOK, it)
}

func DeleteBudget(c echo.Context) error {
    ID := c.Param("id")
    result := txtdb.DeleteBudget(ID)
    return c.JSON(http.StatusOK, result)
}