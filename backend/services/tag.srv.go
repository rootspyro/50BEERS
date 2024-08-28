package services

import (
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/db/repositories"
)

type TagSrv struct {
	repo *repositories.TagRepo
}

func NewTagSrv(repo *repositories.TagRepo) *TagSrv {
	return &TagSrv{
		repo: repo,
	}
}

func(s *TagSrv) GetAllTags() ([]Tag, error) {
	
	data, err := s.repo.GetAllTags()

	if err != nil {
		return nil, err
	}

	var tags []Tag

	for _, tag := range data {
		tags = append(tags, parseTag(tag))
	}

	return tags, nil
}

func parseTag(data models.Tag) Tag {
	return Tag{
		ID: ParsePublicId(data.Name),
		Name: data.Name,
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
