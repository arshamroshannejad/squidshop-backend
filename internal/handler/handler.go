package handler

import (
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/go-playground/validator/v10"
)

type handlerImpl struct {
	userHandler domain.UserHandler
}

func NewHandler(services domain.Service) domain.Handler {
	v := validator.New()
	_ = helper.RegisterValidations(v)
	return &handlerImpl{
		userHandler: NewUserHandler(services, v),
	}
}

func (h *handlerImpl) User() domain.UserHandler {
	return h.userHandler
}
