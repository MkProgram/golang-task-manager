package tasks

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Store interface {
	List() []TaskDTO
	Create(name string) TaskDTO
}

func Routes() http.Handler {
	mux := http.NewServeMux()
	store := NewSQLStore()
	mux.Handle("/task", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, store.List(), http.StatusOK)
		case http.MethodPost:
			var req struct {
				Description string `json:"description"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil || strings.TrimSpace(req.Description) == "" {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			p := store.Create(req.Description)
			writeJSON(w, p, http.StatusCreated)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	return mux
}

func writeJSON(w http.ResponseWriter, v any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
