package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	usersrepository "github.com/Andromeeeda/Subtracker/internal/features/users/repository/users_rep"
	"github.com/jackc/pgx/v5"
)

func (r *UsersPostgresRep) GetByUserID(ctx context.Context, userID string) (*usersrepository.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	SELECT user_id,username,hashPassword FROM subtracker.users
	WHERE user_id = $1
	`

	var u usersrepository.User
	err := r.pool.QueryRow(ctx, query, userID).
		Scan(&u.User_id, &u.Username, &u.HashPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", common_errors.ErrNotFound)
		}
		return nil, fmt.Errorf("find user by id: %w", err)
	}
	return &u, nil

}
