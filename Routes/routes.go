package routes

import (
	m "APILearning/Models"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// use a single instance of Validate, it caches struct info
var uni *ut.UniversalTranslator
var validate *validator.Validate

func SetupRoutes() {

	e := echo.New()

	e.GET("/", helloWorld)
	e.POST("users", createUser)

	e.GET("all", getUser)
	e.GET("beers", getBeers)
	e.GET("beers/:id", getBeer)

	setupValidator()

	e.Logger.Fatal(e.Start(":1323"))

}

func setupValidator() {

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

}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {

	articles := m.Articles{
		m.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		m.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")

	return c.JSON(http.StatusOK, articles)

}

func getBeers(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}

func getBeer(c echo.Context) error {

	id := c.Param("id")
	return c.String(http.StatusOK, id)

}

func createUser(c echo.Context) error {

	u := new(m.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Ocorreu um erro")
	}

	err := validate.Struct(u)
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

	return c.JSON(http.StatusCreated, u)

}
