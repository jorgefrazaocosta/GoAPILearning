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

	g := e.Group("v1/")

	setupRoutesWithoutAccessToken(g)
	setupRouterAccessTokenRequired(g)

	e.Logger.Fatal(e.Start(":1323"))

}

func setupRoutesWithoutAccessToken(g *echo.Group) {

	g.POST("signup", session.SignUp)
	g.POST("signin", session.SignIn)
	g.POST("recover-password", session.RecoverPassword)

}

func setupRouterAccessTokenRequired(g *echo.Group) {

	g.GET("user", user.GetUser, authentication.CustomJWT())
	g.POST("user", user.CreateUser, authentication.CustomJWT())
	g.PUT("user", user.UpdateUser, authentication.CustomJWT())
	g.GET("beer/:id", beer.GetBeer, authentication.CustomJWT())
	g.POST("beer", beer.CreateBeer, authentication.CustomJWT())

}
