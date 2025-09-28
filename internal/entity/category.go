package entity

type CategoryCreateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100" example:"Electronics"`
	Slug     string `json:"slug" validate:"required,min=1,max=100" example:"electronics"`
	ParentID *int   `json:"parent_id,omitempty" validate:"omitempty,min=1,numeric" example:"1"`
}

type CategoryUpdateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100" example:"Electronics"`
	Slug     string `json:"slug" validate:"required,min=1,max=100" example:"electronics"`
	ParentID *int   `json:"parent_id,omitempty" validate:"omitempty,min=1,numeric" example:"1"`
}

type CategoryQueryParamRequest struct {
	Name string `json:"name" validate:"omitempty,min=1,max=100,required_without=Slug" example:"Electronics"`
	Slug string `json:"slug" validate:"omitempty,min=1,max=100,required_without=Name" example:"electronics"`
}
