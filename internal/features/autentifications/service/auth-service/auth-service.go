package authservice

import (
	"errors"

	sessionstorage "github.com/Andromeeeda/Subtracker/internal/features/autentifications/sessionStorage"
	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
	"github.com/Andromeeeda/Subtracker/pkg/HashPassword"
	"github.com/google/uuid"
)

type AuthService struct {
	usersRep *usersrepository.UsersRep
	sessionStore *sessionstorage.SessionStorage
}

func NewAuthService(usersRep *usersrepository.UsersRep,sessionStore *sessionstorage.SessionStorage) *AuthService {
	return &AuthService{
		usersRep: usersRep,
		sessionStore: sessionStore,
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

func (s* AuthService) Login(username string,userPassword string) (string,error){
		user, err := s.usersRep.FindByUSersName(username); 
		if err != nil {
			return "",err 
		}

		if err := hashpassword.CheckHashPassword(user.HashPassword,userPassword); err != nil {
			return "",err
		}

		session_id := s.sessionStore.Create(user.User_id)

		return session_id,nil 
}
