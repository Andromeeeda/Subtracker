package subscription_repository_postgres

import (
	"context"
	"fmt"

	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
)

func (r *SubPostgresRep) GetAllSub(ctx context.Context, userID string) ([]*subrep.SubscriptionModel, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Optimeout())
	defer cancel()

	query := `
	SELECT sub_id,user_id,title,cost,date
	FROM subtracker.subscriptions
	WHERE user_id = $1
	ORDER BY sub_id
	`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query subscriptions: %w", err)
	}
	defer rows.Close()

	var subs []*subrep.SubscriptionModel
	for rows.Next() {
		var m subrep.SubscriptionModel
		if err := rows.Scan(&m.ID, &m.User_id, &m.Title, &m.Cost, &m.Date); err != nil {
			return nil, fmt.Errorf("scan subscription: %w", err)
		}
		subs = append(subs, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate subscriptions: %w", err)
	}

	return subs, nil
}
