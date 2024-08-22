package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db"
	"github.com/rootspyro/50BEERS/db/repositories"
	"github.com/rootspyro/50BEERS/handlers/drinks"
	"github.com/rootspyro/50BEERS/handlers/health"
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
	countriesRepo := repositories.NewCountriesRepo(database.Collection("country"))
	locationRepo := repositories.NewLocationRepo(database.Collection("location"))
	drinksRepo := repositories.NewDrinksRepo(database.Collection("drink"))

	// services
	drinkSrv := services.NewDrinkSrv(countriesRepo, locationRepo, drinksRepo)

	// handlers 
	healthHandler := health.NewHealthHandler()
	drinkHandler := drinks.NewDrinkHandler(drinkSrv)

	// routes
	routes := routes.New(
		healthHandler,
		drinkHandler,
	)
	
	// Configurate server
	app := http.Server{
		Handler: routes,
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
