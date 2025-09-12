package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type logEntry struct {
	Method   string `json:"method"`
	Path     string `json:"path"`
	Status   int    `json:"status"`
	Duration string `json:"duration"`
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)
		entry := logEntry{
			Method:   r.Method,
			Path:     r.URL.Path,
			Status:   lrw.statusCode,
			Duration: time.Since(start).String(),
		}
		b, err := json.Marshal(entry)
		if err != nil {
			log.Printf("error marshaling log: %v", err)
			return
		}
		log.Println(string(b))
	})
}
