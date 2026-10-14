package userservice

import (
	"context"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
	hashpassword "github.com/Andromeeeda/Subtracker/pkg/HashPassword"
)

type UserRep interface {
	Create(ctx context.Context, user *usersrepository.User) error
	FindByUSersName(ctx context.Context, username string) (*usersrepository.User, error)
	GetByUserID(ctx context.Context, userID string) (*usersrepository.User, error)
	ChangeUserPassword(ctx context.Context, userID string, newPassword string) error
	DeleteUser(ctx context.Context, userID string) error
}

type SubRep interface {
	DeleteAllSubByUser(ctx context.Context,userID string) error
}

type UserService struct {
	userRep  UserRep
	sessions *sessionstorage.SessionStorage
	subRep   SubRep
}

func NewUserService(userRep UserRep, sessions *sessionstorage.SessionStorage, subRep SubRep) *UserService {
	return &UserService{
		userRep:  userRep,
		sessions: sessions,
		subRep:   subRep,
	}
}

func (s *UserService) GetUser(ctx context.Context, userID string) (string, error) {
	user, err := s.userRep.GetByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (s *UserService) ChangeUserPassword(ctx context.Context, userID string, currentPassword string, newPassword string) error {

	user, err := s.userRep.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if err := hashpassword.CheckHashPassword(user.HashPassword, currentPassword); err != nil {
		return err
	}

	newHashPassword, err := hashpassword.Hash(newPassword)
	if err != nil {
		return err
	}

	if err := s.userRep.ChangeUserPassword(ctx, userID, newHashPassword); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context,userID string, password string) error {

	user, err := s.userRep.GetByUserID(ctx,userID)
	if err != nil {
		return err
	}

	if err := hashpassword.CheckHashPassword(user.HashPassword, password); err != nil {
		return fmt.Errorf("hashPassword error: %w",common_errors.ErrBadRequest)
	}

	if err := s.sessions.DeleteSession(userID); err != nil {
		return err
	}

	s.subRep.DeleteAllSubByUser(ctx,userID)

	s.userRep.DeleteUser(ctx,userID)

	return nil
}
