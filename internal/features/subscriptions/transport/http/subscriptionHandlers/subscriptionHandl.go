package subscriptionhandlers

import "github.com/Andromeeeda/Subtracker/internal/core/domain"

type SubService interface {
	CreateSubscription(title string,cost float64,date int) (*domain.Subscription,error)
	GetAllSub() []*domain.Subscription
	GetSub(int) (*domain.Subscription,error)
}

type SubHandler struct {
	subService SubService
}

func NewSubHandler(subService SubService) *SubHandler {
	return &SubHandler{
		subService: subService,
	}
}
