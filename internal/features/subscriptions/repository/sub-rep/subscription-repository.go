package subrep

import (
	"errors"

	"github.com/Andromeeeda/Subtracker/internal/core/domain"
)

type SubRepository struct {
	subscriptions map[int]*domain.Subscription
	NextID int 
}

func NewSubRep() *SubRepository {
	return &SubRepository{
		subscriptions: make(map[int]*domain.Subscription),
		NextID: 0,
	}
}

func (r *SubRepository) AddSubscription(title string, cost float64, date int) (*domain.Subscription,error) {
	
	//Проверка на уникальность по названию 
	for _,s := range r.subscriptions{
		if s.Title == title {
			return nil,errors.New("subscription with this title already exist!")
		}
	}

	r.NextID++

	sub := &domain.Subscription{
		ID:    r.NextID,
		Title: title,
		Cost:  cost,
		Date:  date,
	}

	r.subscriptions[r.NextID] = sub

	return sub,nil
}
