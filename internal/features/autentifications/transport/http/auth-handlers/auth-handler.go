package authhandlers

import (
	authservice "github.com/Andromeeeda/Subtracker/internal/features/autentifications/service/auth-service"
)

type AuthHAndler struct {
	authService *authservice.AuthService
}

func NewAuthHandler(authService *authservice.AuthService) *AuthHAndler {
	return &AuthHAndler{
		authService: authService,
	}
}
