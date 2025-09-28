package entity

type ProductRatingRequest struct {
	Rate int `json:"rate" validate:"required,oneof=1 2 3 4 5,numeric" example:"5"`
}
