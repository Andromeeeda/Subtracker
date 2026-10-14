package subscription_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
	"github.com/jackc/pgx/v5"
)

func (r *SubPostgresRep) GetSubscription(ctx context.Context, userID string, id int) (*subrep.SubscriptionModel, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	SELECT sub_id,user_id,title,cost,date 
	FROM subtracker.subscriptions
	WHERE user_id = $1 AND sub_id = $2
	RETURNING sub_id,user_id,title,cost,date 
	`

	row := r.pool.QueryRow(ctx, query, userID, id)

	var subModel subrep.SubscriptionModel
	err := row.Scan(&subModel.ID, &subModel.User_id, &subModel.Title, &subModel.Cost, &subModel.Date)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("subscription not found: %w", common_errors.ErrNotFound)
		} else {
			return nil, fmt.Errorf("change subscription: %w", err)
		}
	}

	return &subModel,nil 
}
