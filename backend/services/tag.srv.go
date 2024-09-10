package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TagSrv struct {
	repo *repositories.TagRepo
}

func NewTagSrv(repo *repositories.TagRepo) *TagSrv {
	return &TagSrv{
		repo: repo,
	}
}

func(s *TagSrv) GetAllTags(lang string) ([]Tag, error) {
	
	data, err := s.repo.GetAllTags()

	if err != nil {
		return nil, err
	}

	var tags []Tag

	for _, tag := range data {
		tags = append(tags, parseTag(tag, lang)) }

	return tags, nil
}

func parseTag(data models.Tag, lang string) Tag {

	var name string = data.EN.Name

	if (lang == "es") {
		name = data.ES.Name
	}

	return Tag{
		ID: ParsePublicId(data.EN.Name),
		Name: cases.Title(language.Und).String(name),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

type Tag struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"UpdatedAt"`
}
