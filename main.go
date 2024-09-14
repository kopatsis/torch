package main

import (
	"log"
	"net/http"
	"torch/config"
	"torch/routing"
)

func main() {

	config.EnvVariables()

	rtr := routing.New()

	port := config.GetPort()

	if err := http.ListenAndServe(":"+port, rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
