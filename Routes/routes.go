package routes

import (
	"api.beermenu.com/components/config"
	middleware "api.beermenu.com/components/middleware"
	beer "api.beermenu.com/controllers/beer"
	session "api.beermenu.com/controllers/session"
	user "api.beermenu.com/controllers/user"
	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo"
	middlewareEcho "github.com/labstack/echo/middleware"
)

func SetupRoutes() {

	e := echo.New()

	e.Use(middlewareEcho.Logger())
	e.Use(middleware.DefaultProperties())

	g := e.Group("v1/")

	setupRoutesWithoutAccessToken(g)
	setupRouterAccessTokenRequired(g)

	if config.Data.Server.Port == ":443" {

		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
		e.Logger.Fatal(e.StartAutoTLS(config.Data.Server.Port))

	} else {
		e.Logger.Fatal(e.Start(config.Data.Server.Port))
	}

}

func setupRoutesWithoutAccessToken(g *echo.Group) {

	g.POST("signup", session.SignUp)
	g.POST("signin", session.SignIn)
	g.POST("recover-password", session.RecoverPassword)

}

func setupRouterAccessTokenRequired(g *echo.Group) {

	g.GET("user", user.GetUser, middleware.CustomJWT())
	g.POST("user", user.CreateUser, middleware.CustomJWT())
	g.PUT("user", user.UpdateUser, middleware.CustomJWT())
	g.GET("beer/:id", beer.GetBeer, middleware.CustomJWT())
	g.POST("beer", beer.CreateBeer, middleware.CustomJWT())

}
