package main

import (
	"HelloWord/serverConfig"
	"HelloWord/serverConfig/internals/handlers"
	"HelloWord/serverConfig/internals/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load config
	config, err := serverConfig.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config %v", err)
	}
	// Create a nde handler
	handler := handlers.NewHandlers()

	//set up the HTTP Server
	mux := http.NewServeMux()

	// Setup Route
	routes.SetupRoutes(mux, handler)
	//server instance
	serverAddress := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddress,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Service fail %v", err)
	}
}
