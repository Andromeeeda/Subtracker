package subrep

import (
	"errors"

	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
)

type SubRepository struct {
	subscriptions map[int]*subrep.SubscriptionModel
	NextID        int
}

func NewSubRep() *SubRepository {
	return &SubRepository{
		subscriptions: make(map[int]*subrep.SubscriptionModel),
		NextID:        0,
	}
}

func (r *SubRepository) AddSubscription(userID string, title string, cost float64, date int) (*subrep.SubscriptionModel, error) {

	//Проверка на уникальность по названию
	for _, s := range r.subscriptions {
		if s.Title == title && s.User_id == userID {
			return nil, errors.New("subscription with this title already exist!")
		}
	}

	r.NextID++

	sub := &subrep.SubscriptionModel{
		ID:      r.NextID,
		User_id: userID,
		Title:   title,
		Cost:    cost,
		Date:    date,
	}

	r.subscriptions[r.NextID] = sub

	return sub, nil
}

func (r *SubRepository) GetAllSub(userID string) []*subrep.SubscriptionModel {

	result := make([]*subrep.SubscriptionModel, 0, len(r.subscriptions))

	for _, s := range r.subscriptions {
		if s.User_id == userID {
			result = append(result, s)
		}
	}

	return result
}

func (r *SubRepository) GetSubscription(userID string, id int) (*subrep.SubscriptionModel, error) {
	sub, ok := r.subscriptions[id]
	if !ok {
		return nil, errors.New("Not found")
	}

	if sub.User_id != userID {
		return nil, errors.New("Not found")
	}

	return sub, nil
}

func (r *SubRepository) ChangeSub(userID string, id int, title *string, cost *float64, date *int) (*subrep.SubscriptionModel, error) {

	m, ok := r.subscriptions[id]
	if !ok {
		return nil, errors.New("Not found")
	}

	if m.User_id != userID {
		return nil, errors.New("Not found")
	}

	if title != nil {
		m.Title = *title
	}

	if cost != nil {
		m.Cost = *cost
	}

	if date != nil {
		m.Date = *date
	}

	return m, nil
}

func (r *SubRepository) DeleteSub(userID string, id int) error {
	sub, ok := r.subscriptions[id]
	if !ok {
		return errors.New("Not Found")
	}

	if sub.User_id != userID {
		return errors.New("Not Found")
	}

	delete(r.subscriptions, id)

	return nil
}
