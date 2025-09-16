package model

type Category struct {
	ID            string     `json:"id" example:"1"`
	Name          string     `json:"name" example:"Laptop"`
	Slug          string     `json:"slug" example:"laptop"`
	ParentID      *string    `json:"-"`
	SubCategories []Category `json:"sub_categories,omitempty" swaggertype:"array" example:"[{\"id\":\"2\",\"name\":\"Gaming Laptop\",\"slug\":\"gaming-laptop\"}]"`
}
