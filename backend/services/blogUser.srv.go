package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"go.mongodb.org/mongo-driver/bson"
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
	user, err := s.repo.GetUser(bson.D{{"username", username}})
	if err != nil {
		return BlogUser{}, err
	}

	return parseBlogUser(user), err
}

func (s *BlogUserSrv) GetUserByEmail(email string) (BlogUser, error) {
	user, err := s.repo.GetUser(bson.D{{"email", email}})
	if err != nil {
		return BlogUser{}, err
	}

	return parseBlogUser(user), err
}

// this function recieves the "user" variable that could be the username or the email
func (s *BlogUserSrv) GetUserForLogin(user string) (BlogUserWithPass, error){
	filter := bson.D{{
		"$or",
		bson.A{
			bson.D{{"username", user}},	
			bson.D{{"email", user}},	
		},
	},}	

	data, err := s.repo.GetUser(filter)
	if err != nil {
		return BlogUserWithPass{}, err
	}

	return BlogUserWithPass{
		Username: data.Username,
		Email: data.Email,
		Password: data.Password,
		Origin: data.Origin,
	}, nil
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

type BlogUserWithPass struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Origin   string `json:"origin"`
	Password string `json:"password"`
}

type BlogUserDTO struct{
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
