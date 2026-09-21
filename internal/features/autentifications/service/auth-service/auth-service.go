package authservice

import (
	"errors"

	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
	"github.com/Andromeeeda/Subtracker/pkg/HashPassword"
	"github.com/google/uuid"
)

type AuthService struct {
	usersRep *usersrepository.UsersRep
}

func NewAuthService(usersRep *usersrepository.UsersRep) *AuthService {
	return &AuthService{
		usersRep: usersRep,
	}
}

func(s *AuthService) Register(username string,userPassword string) error{

	check := s.usersRep.ExistsByUserName(username)
	if check {
		return errors.New("user already exist")
	}

	HashPassword,err := hashpassword.Hash(userPassword)
	if err != nil {
		return err
	}

	user := &usersrepository.User{
		User_id: uuid.New().String(),
		Username: username,
		HashPassword: HashPassword,
	}

	if err := s.usersRep.Create(user); err != nil {
		return err
	}

	return nil 	 
}
