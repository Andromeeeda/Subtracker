package usersrepository

import (
	"errors"

	hashpassword "github.com/Andromeeeda/Subtracker/pkg/HashPassword"
)

type UsersRep struct {
	users map[string]*User
}

func NewUsersRep() *UsersRep {
	return &UsersRep{
		users: make(map[string]*User),
	}
}

func (r *UsersRep) Create(user *User) error {

	
    if _, exists := r.users[user.Username]; exists {
        return errors.New("username already exists")
    }

    r.users[user.Username] = user
    return nil
	
}

func (r *UsersRep) FindByUSersName(username string) (*User, error) {
	user, ok := r.users[username]
	if !ok {
		return nil, errors.New("Error while receiving user data. Incorrect username")
	}

	return user,nil 
}

func(r *UsersRep) ExistsByUserName(username string) bool {
	_,exists := r.users[username]

	return exists
}

func(r *UsersRep) FindByUserID(userID string) (*User,error) {
	for _, user := range r.users {
		if user.User_id == userID{
			return user,nil
		}
	}

	return nil, errors.New("Not Found")
}

func(r *UsersRep) ChangeUserPassword(userID string,newPassword string) error {

	newHashPassword,err := hashpassword.Hash(newPassword)
	if err != nil {
		return err 
	}

	user,err := r.FindByUserID(userID)
	if err != nil {
		return err 
	}

	user.HashPassword = newHashPassword

	return nil 
}

func(r *UsersRep) DeleteUser(userID string) error {

	for _, user := range r.users {
		if user.User_id == userID{
			delete(r.users,userID)
		}
	}

	return nil 
}