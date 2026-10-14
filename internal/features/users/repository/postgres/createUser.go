package users_repository_postgres

import (
	"context"
	"fmt"

	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
)

func (r *UsersPostgresRep) Create(ctx context.Context, user *usersrepository.User) error {
	ctx,cancel := context.WithTimeout(ctx,r.pool.Optimeout())
	defer cancel()
	
	query := `INSERT INTO subtracker.users (username,hashPassword) 
	VALUES ($1,$2)
	RETURNING user_id,username;`

	row := r.pool.QueryRow(ctx, query, user.Username, user.HashPassword)
	
	var userModel UserModel
	err := row.Scan(&userModel.Username,&userModel.HashPassword)

	if err != nil {
		return fmt.Errorf("scan error: %w",err)
	}

	return nil 
}
