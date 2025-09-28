package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/go-playground/validator/v10"
)

type authHandlerImpl struct {
	service   domain.Service
	validator *validator.Validate
}

func NewAuthHandler(service domain.Service, validator *validator.Validate) domain.AuthHandler {
	return &authHandlerImpl{
		service:   service,
		validator: validator,
	}
}

// AuthUserHandler godoc
//
//	@Summary		auth endpoint (register | login)
//	@Description	if user exists it will log in else register. it also sends otp code to user phone
//	@Accept			json
//	@Produce		json
//	@Tags			Auth
//	@Param			request	body	entity.UserAuthRequest	true	"phone for register or login"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/auth [post]
func (u *authHandlerImpl) AuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody entity.UserAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	if err := u.validator.Struct(reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	if err := u.service.User().CreateUser(r.Context(), &reqBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	otpCode, err := u.service.OTP().Generate(r.Context(), reqBody.Phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	smsMsg := fmt.Sprintf("کد احراز هویت شما : %s\nفروشگاه اینترنتی اسکویید شاپ", otpCode)
	if err := u.service.Sms().Send(smsMsg, reqBody.Phone); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// VerifyAuthUserHandler godoc
//
//	@Summary		verify auth endpoint
//	@Description	verify otp code and return access token
//	@Accept			json
//	@Produce		json
//	@Tags			Auth
//	@Param			request	body	entity.UserVerifyAuthRequest	true	"phone and code for register or login"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/auth/verify [post]
func (u *authHandlerImpl) VerifyAuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody entity.UserVerifyAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	if err := u.validator.Struct(reqBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp, _ := json.Marshal(helper.M{"error": err.Error()})
		w.Write(resp)
		return
	}
	isValid, err := u.service.OTP().Verify(r.Context(), reqBody.Phone, reqBody.Code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf(`{"error": "invalid otp code or expired"}`)))
		return
	}
	user, err := u.service.User().GetUserByPhone(r.Context(), reqBody.Phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token, err := u.service.User().GenerateUserJwtToken(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, token)))
}
