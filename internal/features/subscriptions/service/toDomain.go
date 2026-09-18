package service

import (
	"github.com/Andromeeeda/Subtracker/internal/core/domain"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
)

func toDomain(m *subrep.SubscriptionModel) *domain.Subscription {

	if m == nil {
		return nil 
	}

	return &domain.Subscription{
		ID: m.ID,
		Title: m.Title,
		Cost: m.Cost,
		Date: m.Date,
	}
	
}