package app

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Gunga-D/taskL0/internal/config"
)

type Server struct {
	core         *http.Server
	handler      http.Handler
	address      string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func New(config *config.AppConfig, handler http.Handler) (*Server, error) {
	wrappedAddress := ":" + config.Port

	wrappedReadTimeout, err := strconv.Atoi(config.ReadTimeout)
	if err != nil {
		return nil, err
	}

	var wrappedWriteTimeout int
	wrappedWriteTimeout, err = strconv.Atoi(config.WriteTimeout)
	if err != nil {
		return nil, err
	}

	return &Server{
		handler:      handler,
		address:      wrappedAddress,
		readTimeout:  time.Duration(wrappedReadTimeout),
		writeTimeout: time.Duration(wrappedWriteTimeout),
	}, nil
}

func (s *Server) Run() error {
	s.core = &http.Server{
		Addr:           s.address,
		Handler:        s.handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    s.readTimeout * time.Second,
		WriteTimeout:   s.writeTimeout * time.Second,
	}
	return s.core.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.core.Shutdown(ctx)
}
