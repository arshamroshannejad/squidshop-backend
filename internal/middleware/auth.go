package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/helper"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID  string `json:"user_id"`
	Phone   string `json:"phone"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func RequireAuth(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error":"missing Authorization header"}`))
				return
			}
			parts := strings.Fields(authHeader)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error":"invalid Authorization header format"}`))
				return
			}
			tokenString := parts[1]
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
				return []byte(cfg.App.Secret), nil
			})
			if err != nil {
				switch {
				case errors.Is(err, jwt.ErrTokenExpired):
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(`{"error":"token is expired"}`))
				default:
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(`{"error":"invalid token"}`))
				}
				return
			}
			if !token.Valid {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"error":"invalid token"}`))
				return
			}
			ctx := context.WithValue(r.Context(), helper.CtxUserID, claims.UserID)
			ctx = context.WithValue(ctx, helper.CtxPhone, claims.Phone)
			ctx = context.WithValue(ctx, helper.CtxIsAdmin, claims.IsAdmin)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin, ok := r.Context().Value(helper.CtxIsAdmin).(bool)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !isAdmin {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"error":"admin access required"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
