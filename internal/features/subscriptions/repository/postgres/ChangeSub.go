package subscription_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
	"github.com/jackc/pgx/v5"
)

func (r *SubPostgresRep) ChangeSub(ctx context.Context, userID string, id int, title *string, cost *float64, date *int) (*subrep.SubscriptionModel, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	UPDATE subtracker.subscriptions
	SET title = COALESCE($1, title),
            cost  = COALESCE($2, cost),
            date  = COALESCE($3, date)
			WHERE sub_id = $4 AND user_id = $5
			RETURNING sub_id,user_id,title,cost,date
			`

	row := r.pool.QueryRow(ctx, query, title, cost, date, id, userID)

	var subModel subrep.SubscriptionModel
	err := row.Scan(&subModel.ID, &subModel.User_id, &subModel.Title, &subModel.Cost, &subModel.Date)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("subscription not found: %w", common_errors.ErrNotFound)
		} else {
			return nil,fmt.Errorf("change subscription: %w",err)
		}
	}

	return &subModel,nil 
}
