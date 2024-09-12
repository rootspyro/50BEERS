package services

import "github.com/rootspyro/50BEERS/db/repositories"

type SubscriberSrv struct {
	repo *repositories.SubscriberRepo
}

func NewSubscriberSrv(repo *repositories.SubscriberRepo) *SubscriberSrv {
	return &SubscriberSrv{
		repo: repo,
	}
}
