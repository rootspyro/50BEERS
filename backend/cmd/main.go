package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/db"
	"github.com/rootspyro/50BEERS/routes"
)

func main() {

	// configuration
	config.Init()

	// database connection
	DBClient, err := db.New()

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("Connected to database...")
	defer DBClient.Disconnect(context.TODO())

	// routes
	routes := routes.Init()

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
