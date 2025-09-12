package model

import "time"

type User struct {
	ID        string    `json:"id" example:"1"`
	Phone     string    `json:"phone" example:"+989029266635"`
	CreatedAt time.Time `json:"created_at" example:"2025-09-12T00:12:12.123456789Z"`
}
