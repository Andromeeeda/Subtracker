package subscription_repository_postgres

import (
	"context"
	"fmt"
)

func (r *SubPostgresRep) DeleteAllSubByUser(ctx context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	DELETE FROM subtracker.subscription
	WHERE user_id = $1
	`
	if _,err := r.pool.Exec(ctx,query,userID); err != nil {
		return fmt.Errorf("delete all subscription by user: %w",err)
	}

	return nil 
}
