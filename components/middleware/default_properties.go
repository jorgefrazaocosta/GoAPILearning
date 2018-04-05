package middleware

import (
	"net/http"

	"api.beermenu.com/components/config"
	"api.beermenu.com/components/response"
	"github.com/labstack/echo"
)

func DefaultProperties() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			var bundle string

			if c.Request().Method == "GET" {
				bundle = c.Request().Header.Get("Bundle")
			} else {
				bundle = c.FormValue("bundle")
			}

			// deviceId := c.FormValue("deviceId")
			// locale := c.FormValue("locale")

			if bundle == config.Data.DefaultProperties.Bundle {
				return next(c)
			}

			return response.ErrorKey(c, http.StatusUnauthorized, "Application.Error.Bundle")

		}

	}

}
