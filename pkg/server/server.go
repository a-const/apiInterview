package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	srvr *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.srvr = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
	}
	return s.srvr.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.srvr.Shutdown(ctx)
}
