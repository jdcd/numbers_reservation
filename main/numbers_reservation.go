package main

import (
	"github.com/jdcd/numbers_reservation/internal"
	"log"
	"os"
)

func main() {
	router := internal.SetupRouter(internal.GetRouterDependencies())
	port := os.Getenv("PORT")

	err := router.Run()
	if err != nil {
		log.Fatal("unable to start app in ports ", port, err)
	}
}
