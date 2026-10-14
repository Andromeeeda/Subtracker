package authservice

import (
	"context"

	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
	userservice "github.com/Andromeeeda/Subtracker/internal/features/users/service/user_service"
	hashpassword "github.com/Andromeeeda/Subtracker/pkg/HashPassword"
	"github.com/google/uuid"
)

type AuthService struct {
	usersRep     userservice.UserRep
	sessionStore *sessionstorage.SessionStorage
}

func NewAuthService(usersRep userservice.UserRep, sessionStore *sessionstorage.SessionStorage) *AuthService {
	return &AuthService{
		usersRep:     usersRep,
		sessionStore: sessionStore,
	}
}

func (s *AuthService) Register(ctx context.Context, username string, userPassword string) error {

	HashPassword, err := hashpassword.Hash(userPassword)
	if err != nil {
		return err
	}

	user := &usersrepository.User{
		User_id:      uuid.New().String(),
		Username:     username,
		HashPassword: HashPassword,
	}

	if err := s.usersRep.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, username string, userPassword string) (string, error) {
	user, err := s.usersRep.FindByUSersName(ctx, username)
	if err != nil {
		return "", err
	}

	if err := hashpassword.CheckHashPassword(user.HashPassword, userPassword); err != nil {
		return "", err
	}

	session_id := s.sessionStore.Create(user.User_id)

	return session_id, nil
}
