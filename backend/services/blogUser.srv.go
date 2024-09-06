package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"golang.org/x/crypto/bcrypt"
)

type BlogUserSrv struct {
	repo *repositories.BlogUserRepo
}

func NewBlogUserSrv(repo *repositories.BlogUserRepo) *BlogUserSrv {
	return &BlogUserSrv{
		repo: repo,
	}
}

func (s *BlogUserSrv) GetUserByUsername(username string) (BlogUser, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return BlogUser{}, err
	}

	return parseBlogUser(user), err
}

func (s *BlogUserSrv) GetUserByEmail(email string) (BlogUser, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return BlogUser{}, err
	}

	return parseBlogUser(user), err
}

func (s *BlogUserSrv) NewUserFromSite(data BlogUserDTO) (BlogUser, error){

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return BlogUser{}, err
	}

	result, err := s.repo.CreateUser(models.NewBlogUser{
		Username: data.Username,
		Email: data.Email,
		Password: string(passwordHash),
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
