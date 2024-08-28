package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db"
	"github.com/rootspyro/50BEERS/db/migrations"
	"github.com/rootspyro/50BEERS/db/repositories"
	"github.com/rootspyro/50BEERS/db/seeders"
	"github.com/rootspyro/50BEERS/handlers/country"
	"github.com/rootspyro/50BEERS/handlers/drinks"
	"github.com/rootspyro/50BEERS/handlers/health"
	"github.com/rootspyro/50BEERS/handlers/location"
	"github.com/rootspyro/50BEERS/routes"
	"github.com/rootspyro/50BEERS/services"
)

func main() {
	// flags
	var migrate bool
	var seed string

	flag.BoolVar(&migrate, "migrate", false, "Create the required collections on MongoDB")
	flag.StringVar(&seed, "seed", "", "Insert default data on collection (go run main.go -seed [collection])")

	flag.Parse()

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
	tagRepo := repositories.NewTagRepo(database.Collection("tag"))

	if migrate {

		migrationManager := migrations.NewMigrationManager(database)
		if err := migrationManager.Migrate(); err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		return
	}

	if seed != "" {

		if err := seeders.SeedCollection(
			seed,
			countriesRepo,
			locationRepo,
			tagRepo,
		); err != nil {
			log.Error(err.Error())
		}

		return
	}

	// services
	countrySrv := services.NewCountrySrv(countriesRepo)
	locationSrv := services.NewLocationSrv(locationRepo)
	drinkSrv := services.NewDrinkSrv(countriesRepo, locationRepo, drinksRepo)

	// handlers
	healthHandler := health.NewHealthHandler()
	countryHandler := country.NewCountryHandler(countrySrv)
	locationHandler := location.NewLocationHandler(locationSrv)
	drinkHandler := drinks.NewDrinkHandler(drinkSrv)

	// routes
	routes := routes.New(
		healthHandler,
		countryHandler,
		locationHandler,
		drinkHandler,
	)

	// Configurate server
	app := http.Server{
		Handler: routes,
		Addr:    config.App.Server.Socket,
	}

	fmt.Printf(`
   ...
   | |
   | |     
  /   \    50 BEERS API
 |     |   Listening on %s
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
