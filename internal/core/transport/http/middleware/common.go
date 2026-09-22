package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
)

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

			_, ok := sessionStore.Get(sessionID)
			if !ok {
				respondError(w, http.StatusUnauthorized, "invalid authorization header format")
				return
			}

			next.ServeHTTP(w,r)

		})

	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
