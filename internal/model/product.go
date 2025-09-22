package model

import "time"

type Product struct {
	ID               string    `json:"id" example:"1"`
	Name             string    `json:"name" example:"Call of Duty black ops 4"`
	Slug             string    `json:"slug" example:"call-of-duty-black-ops-4"`
	Description      string    `json:"description" example:"Call of Duty black ops 4 is a first-person shooter game"`
	ShortDescription string    `json:"short_description" example:"Call of Duty black ops 4 is a first-person shooter game"`
	Price            float64   `json:"price" example:"19.99"`
	Quantity         int       `json:"quantity" example:"10"`
	CreatedAt        time.Time `json:"created_at" example:"2025-09-12T00:12:12.123456789Z"`
	UpdatedAt        time.Time `json:"updated_at" example:"2025-09-12T00:12:12.123456789Z"`
	CategoryID       string    `json:"category_id" example:"1"`
}
