package usersrepository

import "errors"

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