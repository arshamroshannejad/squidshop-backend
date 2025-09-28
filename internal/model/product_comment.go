package model

import "time"

type ProductComment struct {
	ID        string    `json:"id" example:"1"`
	ProductID string    `json:"product_id" example:"1"`
	UserID    string    `json:"user_id" example:"1"`
	ParentID  *string   `json:"parent_id,omitempty" example:"1"`
	Comment   string    `json:"comment" example:"comment"`
	CreatedAt time.Time `json:"created_at" example:"2025-09-28T01:20:57+03:30"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-09-28T01:20:57+03:30"`
}
