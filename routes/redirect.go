package routes

import (
	"encoding/json"
	"net/http"

	"go-rest-api/store"

	"github.com/go-chi/chi/v5"
)

func RegisterRedirect(r chi.Router) {
	r.Get("/{code}", func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		url, ok := store.Get(code)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{"status": 404, "error": "Link not found!"})
			return
		}
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})
}
