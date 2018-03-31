package routes

import (
	authentication "api.beermenu.com/components/middleware"
	beer "api.beermenu.com/controllers/beer"
	session "api.beermenu.com/controllers/session"
	user "api.beermenu.com/controllers/user"

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

	e.GET("user", user.GetUser, authentication.CustomJWT())
	e.POST("user", user.CreateUser, authentication.CustomJWT())
	e.PUT("user", user.UpdateUser, authentication.CustomJWT())
	e.GET("beer/:id", beer.GetBeer, authentication.CustomJWT())
	e.POST("beer", beer.CreateBeer, authentication.CustomJWT())

}
