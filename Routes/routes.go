package routes

import (
	"HelloWorldAPI/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func SetupRoutes() {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", helloWorld)
	e.POST("users", createUser)

	e.GET("all", getUser)
	e.GET("beers", getBeers)
	e.GET("beers/:id", getBeer)

	e.Logger.Fatal(e.Start(":1323"))

}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {

	articles := models.Articles{
		models.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		models.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
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

func createUser(c echo.Context) (err error) {

	u := new(models.User)

	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, models.ServerError{Message: "erro"})
	}
	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, models.ServerError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)

}
