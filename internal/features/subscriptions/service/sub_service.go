package service

import (
	"context"
	"fmt"

	"github.com/Andromeeeda/Subtracker/internal/core/domain"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
)

type SubRepository interface {
	AddSubscription(ctx context.Context, userID string, title string, cost float64, date int) (*subrep.SubscriptionModel, error)
	GetAllSub(ctx context.Context, userID string) ([]*subrep.SubscriptionModel,error)
	GetSubscription(ctx context.Context, userID string, id int) (*subrep.SubscriptionModel, error)
	ChangeSub(ctx context.Context, userID string, id int, title *string, cost *float64, date *int) (*subrep.SubscriptionModel, error)
	DeleteSub(ctx context.Context, userID string, id int) error
}

type SubService struct {
	subRepository SubRepository
}

func NewSubService(subRepository SubRepository) *SubService {
	return &SubService{
		subRepository: subRepository,
	}
}

func (s *SubService) CreateSubscription(ctx context.Context, userID string, title string, cost float64, date int) (*domain.Subscription, error) {

	sub, err := s.subRepository.AddSubscription(ctx, userID, title, cost, date)
	if err != nil {
		return nil, fmt.Errorf("err: %w", err)
	}

	return &domain.Subscription{
		ID:      sub.ID,
		User_id: userID,
		Title:   title,
		Cost:    cost,
		Date:    date,
	}, nil

}

func (s *SubService) GetAllSub(ctx context.Context,userID string) ([]*domain.Subscription,error) {

	models,err := s.subRepository.GetAllSub(ctx,userID)
	if err != nil {
		return nil,fmt.Errorf("GetAllSub service error: %w",err)
	}

	result := make([]*domain.Subscription, 0, len(models))

	for _, m := range models {
		result = append(result, &domain.Subscription{
			ID:      m.ID,
			User_id: userID,
			Title:   m.Title,
			Cost:    m.Cost,
			Date:    m.Date,
		})
	}
	return result,nil 

}

func (s *SubService) GetSub(ctx context.Context,userID string, id int) (*domain.Subscription, error) {

	sub, err := s.subRepository.GetSubscription(ctx,userID, id)
	if err != nil {
		return nil, err
	}

	Sub := toDomain(sub)

	return Sub, nil

}

func (s *SubService) ChangeSub(ctx context.Context,userID string, id int, title *string, cost *float64, date *int) (*domain.Subscription, error) {

	ChSub, err := s.subRepository.ChangeSub(ctx,userID, id, title, cost, date)
	if err != nil {
		return nil, err
	}

	ChSubDomain := toDomain(ChSub)

	return ChSubDomain, nil
}

func (s *SubService) DeleteSub(ctx context.Context,userID string, id int) error {
	if err := s.subRepository.DeleteSub(ctx,userID, id); err != nil {
		return err
	}

	return nil
}
