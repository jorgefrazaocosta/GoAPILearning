package models

import (
	"database/sql"
)

type ServerError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type User struct {
	ID          string `json:"userId"`
	Name        string `json:"name" form:"name" query:"name" binding:"required" validate:"required"`
	Email       string `json:"email" form:"email" query:"email" validate:"required,email"`
	Base64Image string `json:"image" form:"image" query:"image"`
}

func (u *User) GetUser(db *sql.DB) error {
	return db.QueryRow("SELECT fullName, email FROM User WHERE userId=?", u.ID).Scan(&u.Name, &u.Email)
}
