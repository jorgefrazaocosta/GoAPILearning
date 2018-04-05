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
			return response.ErrorBadRequestWithKey(c, "User.Error.NotFound")
		}

		return c.JSON(http.StatusBadRequest, err.Error())

	}

	return response.Success(c, u.Cleaned())

}

func CreateUser(c echo.Context) error {

	u := new(m.User)

	if err := c.Bind(u); err != nil {
		return response.ErrorBadRequestWithKey(c, "Application.Error.Unknown")
	}

	if err := validator.ValidateStruct(c, u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if res, _ := upload.Image(u.Base64Image, "users"); res == false {
		return response.ErrorBadRequestWithKey(c, "Upload.Image.Error")
	}

	return response.Success(c, u.Cleaned())

}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Olá")
}
