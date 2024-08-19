package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db"
	"github.com/rootspyro/50BEERS/db/models"
	"github.com/rootspyro/50BEERS/routes"
	"github.com/rootspyro/50BEERS/services"
)

func main() {

	// database
	dbclient, err := db.New()

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	defer dbclient.Disconnect(context.TODO())

	database := dbclient.Database(config.App.Database.Name)

	// models
	drinkModel := models.NewDrinkModel(database.Collection("drink"))

	// services
	drinkSrv := services.NewDrinkSrv(&drinkModel)

	test, err := drinkSrv.GetAllDrinks()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println(test)

	// Configurate server
	app := http.Server{
		Handler: &routes.AppRouter,
		Addr: config.App.Server.Socket,
	}

	fmt.Printf(`
   ...
   | |
   | |     
  /   \    50 BEERS API
 |     |   Listening on %s...
 |     |   
 |     |
 |_____|   By %s %s %s

`, 
 	config.App.Server.Socket,
	config.Colors.Cyan,
	config.App.Author.Name,
	config.Colors.Reset,
 )

	fmt.Printf("\n")

	if err := app.ListenAndServe(); err != nil {
		log.Error(err.Error())
	}
}
