package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main () {
	godotenv.Load()
	
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server is starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
