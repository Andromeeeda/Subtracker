package users_repository_postgres

import (
	"context"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
)

func (r *UsersPostgresRep) ChangeUserPassword(ctx context.Context, userID string, newPassword string) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	UPDATE subtracker.users
	SET hashPassword = $1
	WHERE user_id = $2`

	tag, err := r.pool.Exec(ctx, query, newPassword, userID)
	if err != nil {
		return fmt.Errorf("change user password: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("user not found: %w", common_errors.ErrNotFound)
	}

	return nil
}
