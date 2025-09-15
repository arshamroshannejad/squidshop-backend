package handler

import (
	"encoding/json"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	_ "github.com/arshamroshannejad/squidshop-backend/internal/model"
	"github.com/go-playground/validator/v10"
)

type userHandlerImpl struct {
	service   domain.Service
	validator *validator.Validate
}

func NewUserHandler(service domain.Service, validator *validator.Validate) domain.UserHandler {
	return &userHandlerImpl{
		service:   service,
		validator: validator,
	}
}

// UserProfileHandler godoc
//
//	@Summary		user profile endpoint
//	@Description	get user profile and info based on jwt token
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200	{object}	model.User	"User profile data"
//	@Failure		500
//	@Router			/user/profile [get]
func (h *userHandlerImpl) UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	currentUserID := r.Context().Value(helper.CtxUserID).(string)
	user, err := h.service.User().GetUserByID(r.Context(), currentUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
