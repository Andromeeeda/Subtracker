package userservice

import (
	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
	hashpassword "github.com/Andromeeeda/Subtracker/pkg/HashPassword"
)

type UserRep interface {
	Create(user *usersrepository.User) error
	FindByUSersName(username string) (*usersrepository.User, error)
	ExistsByUserName(username string) bool
	FindByUserID(userID string) (*usersrepository.User, error)
	ChangeUserPassword(userID string, newPassword string) error
	DeleteUser(userID string) error
}

type SubRep interface {
	DeleteAllSubByUser(userID string) error
}

type UserService struct {
	userRep  UserRep
	sessions *sessionstorage.SessionStorage
	subRep SubRep
}

func NewUserService(userRep UserRep, sessions *sessionstorage.SessionStorage,subRep SubRep) *UserService {
	return &UserService{
		userRep:  userRep,
		sessions: sessions,
		subRep: subRep,
	}
}

func (s *UserService) GetUser(userID string) (string, error) {
	user, err := s.userRep.FindByUserID(userID)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (s *UserService) ChangeUserPassword(userID string, currentPassword string, newPassword string) error {

	user, err := s.userRep.FindByUserID(userID)
	if err != nil {
		return err
	}

	if err := hashpassword.CheckHashPassword(user.HashPassword, currentPassword); err != nil {
		return err
	}

	if err := s.userRep.ChangeUserPassword(userID, newPassword); err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(userID string, password string) error {

	user, err := s.userRep.FindByUserID(userID)
	if err != nil {
		return err
	}

	if err := hashpassword.CheckHashPassword(user.HashPassword, password); err != nil {
		return err
	}

	
	if err := s.sessions.DeleteSession(userID); err != nil {
		return err
	}

	s.subRep.DeleteAllSubByUser(userID)
	
	s.userRep.DeleteUser(userID)

	return nil
}
