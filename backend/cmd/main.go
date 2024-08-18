package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/routes"
)

func main() {

	config := config.Init()
	routes := routes.Init()

	app := http.Server{
		Handler: routes,
		Addr: config.Server.Socket,
	}

	fmt.Printf(`
   ...
   | |
   | |     
  /   \    50 BEERS API
 |     |   Listening on %s...
 |     |   
 |     |
 |_____|   By %s.


`, 
 	config.Server.Socket,
	config.Author.Name,
 )

	if err := app.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
