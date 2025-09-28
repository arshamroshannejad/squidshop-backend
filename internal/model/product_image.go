package model

type ProductImage struct {
	ID       string `json:"id" example:"1"`
	ImageURL string `json:"image_url" example:"https://example.com/image.jpg"`
	IsMain   bool   `json:"is_main" example:"true"`
}
