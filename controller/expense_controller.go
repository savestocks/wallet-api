package controller

import (
    "net/http"

    "github.com/andersonlira/wallet-api/gateway/txtdb"
    "github.com/andersonlira/wallet-api/usecase"
    "github.com/andersonlira/wallet-api/domain"
	"github.com/labstack/echo/v4"

)


//GetExpenseList return all objects 
func GetExpenseList(c echo.Context) error {

    list := txtdb.GetExpenseList()

	return c.JSON(http.StatusOK, list)
}

func GetExpenseByID(c echo.Context) error {
    ID := c.Param("id")
    it, err := txtdb.GetExpenseByID(ID)
    if err != nil {
        return c.JSON(http.StatusNotFound,it)
    }
    return c.JSON(http.StatusOK, it)
}

func SaveExpense(c echo.Context) error {
    it := domain.Expense{}
    c.Bind(&it)
    expense,err  := usecase.PrepareExpense(it)
    if err != nil {
        return c.JSON(http.StatusNotFound, false)
    }
    return c.JSON(http.StatusCreated, expense)
}

func UpdateExpense(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Expense{}
    c.Bind(&it)
    it = txtdb.UpdateExpense(ID,it)
    return c.JSON(http.StatusOK, it)
}

func DeleteExpense(c echo.Context) error {
    ID := c.Param("id")
    result := txtdb.DeleteExpense(ID)
    return c.JSON(http.StatusOK, result)
}

