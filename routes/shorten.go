package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"go-rest-api/store"

	"github.com/go-chi/chi/v5"
)

func RegisterShorten(r chi.Router) {
	r.Post("/shorten", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.URL == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]any{"status": 400, "error": "URL is required!"})
			return
		}

		code, err := store.Save(body.URL)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"status": 500, "error": "Could not save URL!"})
			return
		}
		host := os.Getenv("BASE_URL")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"short_url": host + "/" + code,
			"code":      code,
		})
	})
}
