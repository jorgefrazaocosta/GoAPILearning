package localization

import (
	"fmt"
	"log"

	m "api-learning/models"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"

	validator "gopkg.in/go-playground/validator.v9"
)

func init() {

	log.Println("Localization Package")

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

}

var uni *ut.UniversalTranslator
var validate *validator.Validate

/*
func SetupLocalization() *ut.Translator {

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

		return trans

}*/

func T(language string, key interface{}, params ...string) string {

	fmt.Printf("key = %s\n", key)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator(language)

	/*t, _ := trans.C("days-left", 2, 0, "2")
	fmt.Println(t)*/

	translation, _ := trans.T(key, params...)

	return translation

}

func ValidateStruct(c echo.Context, s interface{}) []m.ServerError {

	err := validate.Struct(s)
	if err != nil {

		var errors []m.ServerError

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {

			fmt.Println(err)
			errors = append(errors, m.ServerError{Message: "Erro"})
			return errors

		}

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

		return errors

	}

	return nil

}
