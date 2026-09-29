package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
	"go.uber.org/zap"
)

type ctxKey string

const userIDKey ctxKey = "userID"

func Auth(sessionStore *sessionstorage.SessionStorage) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				respondError(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			sessionID := strings.TrimPrefix(authHeader, "Bearer ")
			if sessionID == authHeader {
				respondError(w, http.StatusUnauthorized, "invalid authorization header format")
				return
			}

			userID, ok := sessionStore.Get(sessionID)
			if !ok {
				respondError(w, http.StatusUnauthorized, "invalid authorization header format")
				return
			}
			ctx := context.WithValue(r.Context(), userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))

		})

	}
}

func Logging(logger *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrapped := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(wrapped, r)

			logger.Infow("http request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", wrapped.statusCode,
				"duration", time.Since(start),
			)
		})
	}
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterWrapper) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDKey).(string)
	return id, ok
}

func respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
