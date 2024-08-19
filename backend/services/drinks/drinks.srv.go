package drinks

func GetAllDrinks() ([]Drink, error) {

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
