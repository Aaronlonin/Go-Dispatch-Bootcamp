package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Aaronlonin/capstone/controller"
	"github.com/Aaronlonin/capstone/router"
	"github.com/Aaronlonin/capstone/service"
	"github.com/Aaronlonin/capstone/usecase"
	"github.com/gorilla/handlers"
)

func main() {
	// create instances for the service, usecase, controller and router
	// injecting the corresponding dependencies to each one of them
	pokemonService := service.New(nil)
	pokemonUsecase := usecase.New(pokemonService)
	pokemonController := controller.New(pokemonUsecase)
	httpRouter := router.Setup(pokemonController)

	// Info to set up the server
	// don't use magic naming and magic numbers, there are better ways to do so (viper - covered in another workshop)
	host := "localhost"
	port := 8080

	// create http.Server instance
	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", host, port),
		Handler:           handlers.LoggingHandler(os.Stdout, httpRouter),
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Printf("starting server in address, %s\n", server.Addr)
	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("starting server: %v", err)
	}
}
