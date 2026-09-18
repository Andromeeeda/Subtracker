package service

import (
	"github.com/Andromeeeda/Subtracker/internal/core/domain"
	subrep "github.com/Andromeeeda/Subtracker/internal/features/subscriptions/repository"
)

type SubRepository interface {
	AddSubscription(title string, cost float64, date int) (*subrep.SubscriptionModel, error)
	GetAllSub() []*subrep.SubscriptionModel
	GetSubscription(int) (*subrep.SubscriptionModel,error)
	ChangeSub(id int, title*string,cost *float64,date *int) (*subrep.SubscriptionModel,error)
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

	sub, err := s.subRepository.AddSubscription(title, cost, date)
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

func (s *SubService) GetAllSub() []*domain.Subscription {

	models := s.subRepository.GetAllSub()

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

func (s* SubService) GetSub(id int) (*domain.Subscription,error) {

	sub,err := s.subRepository.GetSubscription(id)
	if err != nil {
		return nil,err
	}

	Sub := toDomain(sub)

	return Sub,nil 

}

func(s *SubService) ChangeSub(id int,title *string,cost *float64,date *int) (*domain.Subscription,error){

	ChSub,err :=  s.subRepository.ChangeSub(id,title,cost,date)
	if err != nil {
		return nil,err 
	}

	ChSubDomain := toDomain(ChSub)

	return ChSubDomain,nil 
}
