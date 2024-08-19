package main

import (
	"fmt"
	"net/http"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/routes"
)

func main() {

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
