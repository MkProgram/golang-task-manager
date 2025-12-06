package server

import (
	"context"
	"net/http"
	"time"

	"mkassel.com/task-manager/internal/api/tasks"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string) *Server {
	s := &http.Server{
		Addr:         addr,
		Handler:      tasks.Routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return &Server{httpServer: s}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
