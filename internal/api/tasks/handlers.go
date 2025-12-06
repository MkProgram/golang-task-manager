package tasks

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type Store interface {
	List(ctx context.Context) ([]TaskDTO, error)
	Create(ctx context.Context, name string) (TaskDTO, error)
}

func Routes(store Store) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/task", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			tasks, err := store.List(ctx)
			if err != nil {
				log.Fatal("get tasks: %w", err)
			}
			writeJSON(w, tasks, http.StatusOK)
		case http.MethodPost:
			var req struct {
				Description string `json:"description"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil || strings.TrimSpace(req.Description) == "" {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			p, err := store.Create(ctx, req.Description)
			if err != nil {
				log.Fatal("create task: %w", err)
			}
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
