package middleware

import (
	"encoding/json"
	"net/http"
	"os"
)

func ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-Key")
		w.Header().Set("Content-Type", "application/json")
		if key == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]any{"status": 401, "error": "X-API-Key missing!"})
			return
		}
		if key != os.Getenv("API_KEY") {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]any{"status": 403, "error": "Invalid API key!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
