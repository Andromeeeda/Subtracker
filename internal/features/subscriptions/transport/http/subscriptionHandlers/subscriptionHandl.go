package subscriptionhandlers

import (
	"context"

	"github.com/Andromeeeda/Subtracker/internal/core/domain"
)

type SubService interface {
	CreateSubscription(ctx context.Context, userID string, title string, cost float64, date int) (*domain.Subscription, error)
	GetAllSub(ctx context.Context,userID string) ([]*domain.Subscription,error)
	GetSub(ctx context.Context,userID string, id int) (*domain.Subscription, error)
	ChangeSub(ctx context.Context,userID string, id int, title *string, cost *float64, date *int) (*domain.Subscription, error)
	DeleteSub(ctx context.Context,userID string, id int) error
}

type SubHandler struct {
	subService SubService
}

func NewSubHandler(subService SubService) *SubHandler {
	return &SubHandler{
		subService: subService,
	}
}
