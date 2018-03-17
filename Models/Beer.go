package models

type Beer struct {
	Name    string `json:"name" form:"name" query:"name" binding:"required" validate:"required"`
	Brewery string `json:"brewery" form:"brewery" query:"brewery" binding:"required" validate:"required"`
	Country string `json:"country" form:"country" query:"country"`
}

type Beers []Beer
