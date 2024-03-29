package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func server(apiConfig *Config) {

	// Define CORS options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"}, // You can customize this based on your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // You can customize this based on your needs
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for cache, in seconds
	}
	router := chi.NewRouter()
	apiRoute := chi.NewRouter()
	router.Use(cors.Handler(corsOptions))
	apiRoute.Get("/hello", helloReady)
	apiRoute.Get("/error", errorReady)

	//states handlers
	apiRoute.Get("/states", apiConfig.getStatesHandler)
	apiRoute.Get("/states/{state}", apiConfig.getStatesAHandler)
	apiRoute.Get("/state/{state}", apiConfig.getStateHandler)

	// lgas handlers
	apiRoute.Get("/local-government-areas", apiConfig.getLgasHandler)
	apiRoute.Get("/local-government-areas/{state}", apiConfig.getLgasAHandler)

	// cities handlers
	apiRoute.Get("/cities", apiConfig.getCitiesHandler)
	apiRoute.Get("/cities/{state}", apiConfig.getStateCitiesHandler)

	router.Mount("/api", apiRoute)
	router.Get("/", renderHome)
	srv := &http.Server{
		Addr:              ":" + apiConfig.PORT,
		Handler:           router,
		ReadHeaderTimeout: time.Minute,
	}
	log.Printf("serving server on port %v", apiConfig.PORT)

	log.Printf("Serving on port: %s\n", apiConfig.PORT)
	log.Fatal(srv.ListenAndServe())
}
