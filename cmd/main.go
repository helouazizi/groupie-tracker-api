package main

import (
	"fmt"
	"log"
	"net/http"

	"go-rest-api/internal/api"
	"go-rest-api/internal/dependencies"
	"go-rest-api/pkg/config"
	"go-rest-api/pkg/logger"
)

func main() {
	cg := config.Load()
	// Get the current working directory
	logger, err := logger.Create_Logger(cg.LogPath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer logger.Close()

	deps := dependencies.NewDependencies()
	router := api.NewRouter(deps)

	fmt.Printf("Server started at :%s", cg.Port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
