package handler

import (
	"fmt"
	"net/http"
	"sync"

	"go-rest-api/db"
	"go-rest-api/middleware"
	"go-rest-api/routes"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

var (
	once   sync.Once
	router http.Handler
)

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		godotenv.Load()

		mux := chi.NewRouter()
		mux.Use(chimiddleware.Logger)

		routes.RegisterRoot(mux)

		if err := db.Connect(); err != nil {
			fmt.Println("DB-fel:", err)
		} else {
			routes.RegisterRedirect(mux)

			mux.Group(func(mux chi.Router) {
				mux.Use(middleware.ApiKey)
				routes.RegisterShorten(mux)
				routes.RegisterDelete(mux)
			})
		}

		router = mux
	})

	router.ServeHTTP(w, r)
}
