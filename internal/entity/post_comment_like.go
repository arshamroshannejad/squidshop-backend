package entity

type PostCommentLikeCreateUpdate struct {
	Vote int `json:"vote" validate:"required,oneof=1 -1,numeric" example:"1"`
}
