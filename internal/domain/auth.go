package domain

import "net/http"

type AuthHandler interface {
	AuthUserHandler(w http.ResponseWriter, r *http.Request)
	VerifyAuthUserHandler(w http.ResponseWriter, r *http.Request)
}
