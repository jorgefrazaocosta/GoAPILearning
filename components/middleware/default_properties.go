package middleware

import (
	"api.beermenu.com/components/config"
	"github.com/labstack/echo"
)

func DefaultProperties() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			bundle := c.FormValue("bundle")
			// deviceId := c.FormValue("deviceId")
			// locale := c.FormValue("locale")

			if bundle == config.Data.DefaultProperties.Bundle {
				return next(c)
			}

			return &echo.HTTPError{
				Code:    401,
				Message: "É necessário um bundle válido",
			}

		}

	}

}
