package main

import (
	"log"
	"net/http"
	"torch/config"
	"torch/routing"
)

func main() {

	config.EnvVariables()

	// mysql, redis := config.NewMySQLConnection(), config.NewRedisClient()

	// fullService := data.NewMainService(mysql, redis)

	// rtr := routing.New(fullService)

	rtr := routing.TestNew()

	port := config.GetPort()

	if err := http.ListenAndServe(":"+port, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
