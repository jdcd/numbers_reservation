package main

import (
	"fmt"
	"github.com/jdcd/numbers_reservation/internal"
	"github.com/jdcd/numbers_reservation/pkg"
	"os"
)

func main() {
	router := internal.SetupRouter(internal.GetRouterDependencies())
	port := os.Getenv("PORT")

	err := router.Run()
	if err != nil {
		errorDetail := fmt.Sprintf("unable to start app on the port: %s , %s", port, err.Error())
		pkg.ErrorLogger().Fatal(errorDetail)
	}
}
