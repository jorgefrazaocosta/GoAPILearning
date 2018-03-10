package validator

import (
	"fmt"
	"log"
	"net/http"

	m "api-learning/models"

	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt"
	validator "gopkg.in/go-playground/validator.v9"

	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func init() {
	log.Println("<package_name> - Two")
}

// use a single instance of Validate, it caches struct info
var uni *ut.UniversalTranslator
var validate *validator.Validate

func SetupValidator() *validator.Validate {

	en := en.New()
	uni = ut.New(en, en, pt.New())

	err := uni.Import(ut.FormatJSON, "translations")
	if err != nil {
		log.Fatal(err)
	}

	err = uni.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("pt")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return validate

}

func ValidateStruct(c echo.Context, s interface{}) error {

	err := validate.Struct(s)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return c.JSON(http.StatusBadRequest, "Erro")
		}

		var errors []m.ServerError

		trans, _ := uni.GetTranslator("pt")
		t, _ := trans.C("days-left", 2, 0, "2")
		fmt.Println(t)

		for _, err := range err.(validator.ValidationErrors) {

			errors = append(errors, m.ServerError{Message: err.Translate(trans)})

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

		}

		return c.JSON(http.StatusBadRequest, errors)
	}

	return nil

}
