package services 

import "github.com/rootspyro/50BEERS/db/models"

type DrinkSrv struct {
	model *models.DrinkModel
}

func NewDrinkSrv(model *models.DrinkModel) *DrinkSrv {
	return &DrinkSrv{
		model: model,
	}
}


func(s *DrinkSrv) GetAllDrinks() ([]Drink, error) {

	var tempData []Drink = []Drink{
		{
			ID: "tempID-01",
			Name: "Estrella Damn clasica",
			Type: "Pilsner",
		},
		{
			ID: "tempID-02",
			Name: "Voll Damn",
			Type: "double malt",
		},
	}

	return tempData, nil
}

type Drink struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
