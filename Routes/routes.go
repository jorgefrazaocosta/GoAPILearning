package routes

import (
	beer "api-learning/controllers/beer"
	session "api-learning/controllers/session"
	user "api-learning/controllers/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetupRoutes() {

	e := echo.New()
	e.Use(middleware.Logger())

	setupRoutesWithoutAccessToken(e)
	setupRouterAccessTokenRequired(e)

	e.Logger.Fatal(e.Start(":1323"))

}

func setupRoutesWithoutAccessToken(e *echo.Echo) {

	e.POST("signup", session.SignUp)
	e.POST("signin", session.SignIn)
	e.POST("recover-password", session.RecoverPassword)

}

func setupRouterAccessTokenRequired(e *echo.Echo) {

	e.GET("user", user.GetUser)
	e.POST("user", user.CreateUser)
	e.PUT("user", user.UpdateUser)
	e.GET("beer/:id", beer.GetBeer)
	e.POST("beer", beer.CreateBeer)

}
