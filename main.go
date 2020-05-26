package main

import (
	"errors"
	"github.com/gorilla/handlers"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"log"
	"multi-event-booking/routes"
	"net/http"
	"os"
)

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	port, portExists := os.LookupEnv("PORT")

	if !portExists {
		port = "8080"
	}


	var handler http.Handler
	{
    	handler = handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "Authorization", "Content-Type"}),
			handlers.ExposedHeaders([]string{""}),
			handlers.MaxAge(10),
			handlers.AllowCredentials(),
    	)(routes.Init())
    	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
	}

	// http.Handle("/", handler)
	if err := http.ListenAndServe(":"+port, routes.Init()); err != nil {
		logger.Log(err)
		log.Fatal(err)
	}
}