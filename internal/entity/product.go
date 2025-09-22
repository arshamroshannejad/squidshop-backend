package entity

type ProductCreateRequest struct {
	Name             string  `json:"name" validate:"required,min=1,max=255" example:"call of duty black ops 4"`
	Slug             string  `json:"slug" validate:"required,min=1,max=255" example:"call-of-duty-black-ops-4"`
	Description      string  `json:"description" validate:"required,min=1" example:"lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	ShortDescription string  `json:"short_description" validate:"required,min=1,max=255" example:"lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	Price            float64 `json:"price" validate:"required,min=1" example:"23400.23"`
	Quantity         int     `json:"quantity" validate:"required,numeric,min=1" example:"10"`
	CategoryID       int     `json:"category_id" validate:"required,numeric,min=1" example:"1"`
}

type ProductUpdateRequest struct {
	Name             string  `json:"name" validate:"required,min=1,max=255" example:"call of duty black ops 4"`
	Slug             string  `json:"slug" validate:"required,min=1,max=255" example:"call-of-duty-black-ops-4"`
	Description      string  `json:"description" validate:"required,min=1" example:"lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	ShortDescription string  `json:"short_description" validate:"required,min=1,max=255" example:"lorem ipsum dolor sit amet, consectetur adipiscing elit"`
	Price            float64 `json:"price" validate:"required,min=1" example:"23400.23"`
	Quantity         int     `json:"quantity" validate:"required,numeric,min=1" example:"10"`
	CategoryID       int     `json:"category_id" validate:"required,numeric,min=1" example:"1"`
}
