package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Port not found")
		return
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5 minutes
	}))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Mount("/v1", v1Router)

	// Server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	fmt.Printf("Server starting at %+v", PORT)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
