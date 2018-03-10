package Session

import (
	"net/http"

	"github.com/labstack/echo"
)

func SignUp(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}

func SignIn(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}

func RecoverPassword(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}
