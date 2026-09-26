package subscriptionhandlers

import "github.com/Andromeeeda/Subtracker/internal/core/domain"

type SubService interface {
	CreateSubscription(userID string,title string,cost float64,date int) (*domain.Subscription,error)
	GetAllSub(userID string) []*domain.Subscription
	GetSub(userID string,id int) (*domain.Subscription,error)
	ChangeSub(userID string,id int,title *string,cost *float64,date *int) (*domain.Subscription,error)
	DeleteSub(userID string,id int) error
}

type SubHandler struct {
	subService SubService
}

func NewSubHandler(subService SubService) *SubHandler {
	return &SubHandler{
		subService: subService,
	}
}
