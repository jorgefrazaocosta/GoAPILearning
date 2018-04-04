package User

import (
	"database/sql"
	"net/http"

	database "api.beermenu.com/components/database"
	"api.beermenu.com/components/response"
	upload "api.beermenu.com/components/upload"
	validator "api.beermenu.com/components/validator"
	m "api.beermenu.com/models"

	"github.com/labstack/echo"
)

func GetUser(c echo.Context) error {

	u := m.User{ID: "58c080c98a34f"}
	if err := u.GetUser(database.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			return c.JSON(http.StatusBadRequest, "User not found")
		}
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)

}

func CreateUser(c echo.Context) error {

	u := new(m.User)

	if err := c.Bind(u); err != nil {
		return response.ErrorLocalizedKey(c, http.StatusBadRequest, "unknown")
	}

	if err := validator.ValidateStruct(c, u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if res, _ := upload.Image(u.Base64Image, "users"); res == false {
		return c.JSON(http.StatusBadRequest, "Não foi possivel fazer o upload da image")
	}

	return response.Success(c, u.Cleaned())

}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}
