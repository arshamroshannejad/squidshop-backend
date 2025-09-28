package middleware

import (
	"net/http"
	"time"
)

func Timeout(next http.Handler) http.Handler {
	return http.TimeoutHandler(next, 10*time.Second, "Request timeout")
}
