package User

import (
	"net/http"

	validator "api-learning/components/validator"
	m "api-learning/models"

	"github.com/labstack/echo"
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}

func CreateUser(c echo.Context) error {

	u := new(m.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Ocorreu um erro")
	}

	if err := validator.ValidateStruct(c, u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, u)

}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}
