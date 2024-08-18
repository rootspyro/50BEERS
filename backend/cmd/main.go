package main

import (
	"fmt"
	"net/http"

	"github.com/rootspyro/50BEERS/config"
	"github.com/rootspyro/50BEERS/config/log"
	"github.com/rootspyro/50BEERS/routes"
)

func main() {

	sysConf := config.Init()
	routes := routes.Init()

	app := http.Server{
		Handler: routes,
		Addr: sysConf.Server.Socket,
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
 	sysConf.Server.Socket,
	config.Colors.Cyan,
	sysConf.Author.Name,
	config.Colors.Reset,
 )

	fmt.Printf("\n")

	if err := app.ListenAndServe(); err != nil {
		log.Error(err.Error())
	}
}
