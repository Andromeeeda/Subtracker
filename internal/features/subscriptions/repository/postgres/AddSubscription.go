package subscription_repository_postgres

import (
	"context"
	"fmt"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
)

func (r *SubPostgresRep) AddSubscription(ctx context.Context, userID string, title string, cost float64, date int) (*subrep.SubscriptionModel, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	INSERT INTO subtracker.subscriptions (user_id,title,cost,date)
	VALUES ($1,$2,$3,$4)
	RETURNING sub_id,user_id,title,cost,date`

	row := r.pool.QueryRow(ctx, query, userID, title, cost, date)

	var subModel subrep.SubscriptionModel
	err := row.Scan(&subModel.ID, &subModel.User_id, &subModel.Title, &subModel.Cost, &subModel.Date)

	if err != nil {
		return nil,fmt.Errorf("scan error: %w", common_errors.ErrConflict)
	}

	return &subModel,nil 
}
