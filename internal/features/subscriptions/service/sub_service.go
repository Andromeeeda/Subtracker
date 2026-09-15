package service

import (
	"github.com/Andromeeeda/Subtracker/internal/core/domain"
)

type SubRepository interface {
	AddSubscription(title string, cost float64, date int) (*domain.Subscription,error)
}

type SubService struct {
	subRepository SubRepository
}

func NewSubService(subRepository SubRepository) *SubService {
	return &SubService{
		subRepository: subRepository,
	}
}

func (s *SubService) CreateSubscription(title string, cost float64, date int) (*domain.Subscription, error) {

	sub,err := s.subRepository.AddSubscription(title, cost, date)
	if err != nil {
		return nil, err
	}

	return &domain.Subscription{
		ID:    sub.ID,
		Title: title,
		Cost:  cost,
		Date:  date,
	}, nil

}
