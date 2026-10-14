package subscription_repository_postgres

import (
	"context"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
)

func (r *SubPostgresRep) DeleteSub(ctx context.Context, userID string, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `DELETE FROM subtracker.subscriptions
	WHERE sub_id = $1 AND user_id = $2
	`
	tag,err := r.pool.Exec(ctx,query,id,userID)

	if err != nil {
		return fmt.Errorf("delete subscription: %w",err)
	}	
	 if tag.RowsAffected() == 0 {
		return fmt.Errorf("subscription not found: %w",common_errors.ErrNotFound)
	 }

	 return nil 
}
