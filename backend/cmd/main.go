package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rootspyro/50BEERS/config"
)

func main() {

	config := config.Init()

	app := http.Server{
		Addr: config.Server.Socket,
	}

	log.Println("server is starting...")

	fmt.Printf(`
   ...
   | |
   | |
  /   \
 |     |   50 BEERS API
 |     |   Listening on %s...
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
