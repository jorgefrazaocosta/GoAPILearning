package Beer

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetBeer(c echo.Context) error {

	id := c.Param("id")
	return c.String(http.StatusOK, id)

}
