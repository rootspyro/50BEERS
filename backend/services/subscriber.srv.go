package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
)

type SubscriberSrv struct {
	repo *repositories.SubscriberRepo
}

func NewSubscriberSrv(repo *repositories.SubscriberRepo) *SubscriberSrv {
	return &SubscriberSrv{
		repo: repo,
	}
}

func (s *SubscriberSrv) FindByEmail(email string) (Subscriber, error) {

	data, err := s.repo.FindByEmail(email)
	if err != nil {
		return Subscriber{}, err 
	}

	return subscriberParser(data), nil
}

func (s *SubscriberSrv) NewSubsciber(email string) (Subscriber, error) {

	newSubscriber, err := s.repo.NewSubsciber(email)
	if err != nil {
		return Subscriber{}, err
	}
	
	return subscriberParser(newSubscriber), nil
}

func subscriberParser(data models.Subscriber) Subscriber {
	return Subscriber{
		Email: data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

type SubscriberDTO struct {
	Email string `json:"email"`
}

type Subscriber struct {
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
