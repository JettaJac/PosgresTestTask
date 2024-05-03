package server

import (
	"main/internal/config"
	"net/http"
	"time"
	// "main/internal/config"
)

type server struct {
	router       *http.ServeMux
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	// logger       *logrus.Logger
	// storage storage.Storage
}

func NewServer(config *config.Config /*store storage.Storage ,sessionStore sessions.Store*/) *server {
	s := &server{
		router:       http.NewServeMux(),
		Addr:         config.Address,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
		// logger: logrus.New(),
		// storage: Storage,
		// sessionStore: sessionStore,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	// s.router.HandleFunc("/", s.handleHome()).Methods("GET")
	// s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
}
