package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/test/interactive/internal/transport/http"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(port int) *HTTPServer {
	router := transport.NewRouter()
	
	return &HTTPServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
	}
}

func (s *HTTPServer) Start() error {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
