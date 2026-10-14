package usersrepository

import (
	"context"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
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

func (r *UsersRep) Create(ctx context.Context,user *User) error {

	if _, exists := r.users[user.Username]; exists {
		return fmt.Errorf("username already exists %w", common_errors.ErrConflict)
	}

	r.users[user.Username] = user
	return nil

}

func (r *UsersRep) FindByUSersName(ctx,username string) (*User, error) {
	user, ok := r.users[username]
	if !ok {
		return nil, fmt.Errorf("Error while receiving user data. Incorrect username %w", common_errors.ErrBadRequest)
	}

	return user, nil
}

func (r *UsersRep) FindByUserID(userID string) (*User, error) {
	for _, user := range r.users {
		if user.User_id == userID {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found %w", common_errors.ErrNotFound)
}

func (r *UsersRep) ChangeUserPassword(userID string, newPassword string) error {

	newHashPassword, err := hashpassword.Hash(newPassword)
	if err != nil {
		return err
	}

	user, err := r.FindByUserID(userID)
	if err != nil {
		return err
	}

	user.HashPassword = newHashPassword

	return nil
}

func (r *UsersRep) DeleteUser(ctx context.Context,userID string) error {

	for _, user := range r.users {
		if user.User_id == userID {
			delete(r.users, userID)
		}
	}

	return nil
}
