package Beer

import (
	"api-learning/components/validator"
	"fmt"
	"net/http"

	m "api-learning/models"

	"github.com/labstack/echo"
)

func GetBeer(c echo.Context) error {

	id := c.Param("id")
	return c.String(http.StatusOK, id)

}

func CreateBeer(c echo.Context) error {

	fmt.Printf("Accept-Language %s", c.Request().Header.Get("Accept-Language"))

	b := new(m.Beer)

	if err := c.Bind(b); err != nil {
		return c.JSON(http.StatusBadRequest, "Ocorreu um erro")
	}

	if err := validator.ValidateStruct(c, b); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, b)

}