package entity

type UserAuthRequest struct {
	Phone string `json:"phone" validate:"required,min=1,max=13,irphone" example:"+989029266635"`
}

type UserVerifyAuthRequest struct {
	Phone string `json:"phone" validate:"required,irphone" example:"+989029266635"`
	Code  string `json:"code" validate:"required,len=6,numeric" example:"123456"`
}
