package models

type ServerError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type User struct {
	Name  string `json:"name" form:"name" query:"name" binding:"required" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
}
