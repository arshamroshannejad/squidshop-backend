package handler

import (
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/go-playground/validator/v10"
)

type handlerImpl struct {
	authHandler     domain.AuthHandler
	userHandler     domain.UserHandler
	categoryHandler domain.CategoryHandler
	productHandler  domain.ProductHandler
}

func NewHandler(services domain.Service) domain.Handler {
	v := validator.New()
	_ = helper.RegisterValidations(v)
	return &handlerImpl{
		authHandler:     NewAuthHandler(services, v),
		userHandler:     NewUserHandler(services, v),
		categoryHandler: NewCategoryHandler(services, v),
		productHandler:  NewProductHandler(services, v),
	}
}

func (h *handlerImpl) Auth() domain.AuthHandler {
	return h.authHandler
}

func (h *handlerImpl) User() domain.UserHandler {
	return h.userHandler
}

func (h *handlerImpl) Category() domain.CategoryHandler {
	return h.categoryHandler
}

func (h *handlerImpl) Product() domain.ProductHandler {
	return h.productHandler
}
