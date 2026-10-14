package users_repository_postgres

import (
	"context"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
)

func (r *UsersPostgresRep) DeleteUser(ctx context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	DELETE FROM subtracker.users
	WHERE user_id = $1
	`

	tag, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user not found: %w", common_errors.ErrNotFound)
	}

	return nil

}
