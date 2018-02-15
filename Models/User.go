package models

type User struct {
	Name     string `json:"name" form:"name" validate:"required"`
	LastName string `validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
}

type ServerError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// User contains user information
/*type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}*/
