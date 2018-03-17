package validator

import (
	"log"

	"github.com/labstack/echo"

	localize "api-learning/components/localization"
	m "api-learning/models"

	validator "gopkg.in/go-playground/validator.v9"
)

func init() {

	log.Println("Validator Package")

	validate = validator.New()

}

var validate *validator.Validate

func ValidateStruct(c echo.Context, s interface{}) []m.ServerError {

	language := c.Request().Header.Get("Accept-Language")

	err := validate.Struct(s)
	if err != nil {

		var errors []m.ServerError

		if _, ok := err.(*validator.InvalidValidationError); ok {

			errors = append(errors, m.ServerError{Message: localize.T(language, "unknown")})
			return errors

		}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, m.ServerError{Message: localize.T(language, err.Tag(), err.Field())})
		}

		return errors
	}

	return nil

}
