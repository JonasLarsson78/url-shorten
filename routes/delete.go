package routes

import (
	"encoding/json"
	"net/http"

	"go-rest-api/store"

	"github.com/go-chi/chi/v5"
)

func RegisterDelete(r chi.Router) {
	r.Delete("/url/{code}", func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		w.Header().Set("Content-Type", "application/json")

		found, err := store.Delete(code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{"status": 500, "error": "Could not delete URL!"})
			return
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{"status": 404, "error": "Link not found!"})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{"status": 200, "message": "Link deleted!"})
	})
}
