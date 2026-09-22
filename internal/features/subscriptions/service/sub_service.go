package service

import (
	"github.com/Andromeeeda/Subtracker/internal/core/domain"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
)

type SubRepository interface {
	AddSubscription(userID string,title string, cost float64, date int) (*subrep.SubscriptionModel, error)
	GetAllSub(userID string) []*subrep.SubscriptionModel
	GetSubscription(userID string, id int) (*subrep.SubscriptionModel,error)
	ChangeSub(userID string,id int, title*string,cost *float64,date *int) (*subrep.SubscriptionModel,error)
	DeleteSub(userID string,id int) error 
}

type SubService struct {
	subRepository SubRepository
}

func NewSubService(subRepository SubRepository) *SubService {
	return &SubService{
		subRepository: subRepository,
	}
}

func (s *SubService) CreateSubscription(userID string,title string, cost float64, date int) (*domain.Subscription, error) {

	sub, err := s.subRepository.AddSubscription(userID,title, cost, date)
	if err != nil {
		return nil, err
	}

	return &domain.Subscription{
		ID:    sub.ID,
		User_id: userID,
		Title: title,
		Cost:  cost,
		Date:  date,
	}, nil

}

func (s *SubService) GetAllSub(userID string) []*domain.Subscription {

	models := s.subRepository.GetAllSub(userID)

	result := make([]*domain.Subscription,0,len(models))

	for _, m := range models {
		result = append(result, &domain.Subscription{
			ID:    m.ID,
			Title: m.Title,
			Cost:  m.Cost,
			Date:  m.Date,
		})
	}
	return result

}

func (s* SubService) GetSub(userID string,id int) (*domain.Subscription,error) {

	sub,err := s.subRepository.GetSubscription(userID,id)
	if err != nil {
		return nil,err
	}

	Sub := toDomain(sub)

	return Sub,nil 

}

func(s *SubService) ChangeSub(userID string,id int,title *string,cost *float64,date *int) (*domain.Subscription,error){

	ChSub,err :=  s.subRepository.ChangeSub(userID,id,title,cost,date)
	if err != nil {
		return nil,err 
	}

	ChSubDomain := toDomain(ChSub)

	return ChSubDomain,nil 
}

func(s *SubService) DeleteSub(userID string,id int) error {
	if err := s.subRepository.DeleteSub(userID,id); err != nil {
		return err 
	}

	return nil 
}


