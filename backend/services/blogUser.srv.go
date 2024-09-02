package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
)

type BlogUserSrv struct {
	repo *repositories.BlogUserRepo
}

func NewBlogUserSrv(repo *repositories.BlogUserRepo) *BlogUserSrv {
	return &BlogUserSrv{
		repo: repo,
	}
}

func (s *BlogUserSrv) NewUserFromSite(data BlogUserDTO) (BlogUser, error){

	result, err := s.repo.CreateUser(models.NewBlogUser{
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
		Origin: "site",
	})

	if err != nil {
		return BlogUser{}, err
	}

	return parseBlogUser(result), nil 
}

func parseBlogUser(data models.BlogUser) BlogUser {
	return BlogUser{
		Username: data.Username,
		Email: data.Email,
		Origin: data.Origin,
	}
}

type BlogUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Origin   string `json:"origin"`
}

type BlogUserDTO struct{
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
