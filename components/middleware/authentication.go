package middleware

import (
	"net/http"
	"strings"

	database "api.beermenu.com/components/database"

	"github.com/labstack/echo"
)

type (
	JWTConfig struct {
		ContextKey  string
		TokenLookup string
		AuthScheme  string
	}

	jwtExtractor func(echo.Context) (string, error)
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
)

var (
	DefaultJWTConfig = JWTConfig{
		ContextKey:  "user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	}
)

// For invalid token, it returns "401 - Unauthorized" error.
// For missing token, it returns "400 - Bad Request" error.
func CustomJWT() echo.MiddlewareFunc {

	c := DefaultJWTConfig
	return JWTWithConfig(c)

}

func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {

	parts := strings.Split(config.TokenLookup, ":")
	extractor := jwtFromHeader(parts[1], config.AuthScheme)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			auth, err := extractor(c)
			if err != nil {
				return err
			}

			isValid := database.AccessTokenIsValid(auth)

			if err == nil && isValid {

				// Store user information from token into context.
				c.Set(config.ContextKey, "wowowo" /*token*/)
				return next(c)

			}

			return &echo.HTTPError{
				Code:    ErrJWTInvalid.Code,
				Message: ErrJWTInvalid.Message,
				Inner:   err,
			}
		}
	}
}

func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}
