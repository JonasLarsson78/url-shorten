package main

import (
	"fmt"
	"net/http"

	"go-rest-api/db"
	"go-rest-api/middleware"
	"go-rest-api/routes"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := db.Connect(); err != nil {
		fmt.Println("DB-fel:", err)
		return
	}

	port := "8080"
	r := chi.NewRouter()
	r.Use(chimiddleware.Logger)

	routes.RegisterRoot(r)
	routes.RegisterRedirect(r)

	r.Group(func(r chi.Router) {
		r.Use(middleware.ApiKey)
		routes.RegisterShorten(r)
		routes.RegisterDelete(r)
	})

	fmt.Printf("Api is running on http://localhost:%s\n", port)

	http.ListenAndServe(":"+port, r)
}
