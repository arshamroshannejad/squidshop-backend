package entity

type ProductCommentCreateRequest struct {
	Comment  string `json:"comment" validate:"required,min=1" example:"nice product"`
	ParentID *int   `json:"parent_id" example:"12"`
}

type ProductCommentUpdateRequest struct {
	Comment string `json:"comment" validate:"required,min=1" example:"nice product"`
}
