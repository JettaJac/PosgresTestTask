package server

import (
	"encoding/json"
	"log/slog"
	"main/internal/config"
	"main/internal/storage"
	"net/http"
	"time"
)

var (
	PathSave   = "/command/save"   // handleSaveRunCommand
	PathFind   = "/command/find"   //HandleGetOneCommand
	PathList   = "/commands/all"   // HandleGetListCommands
	PathDelete = "/command/delete" // HandleDeleteCommand
)

// server struct
type server struct {
	router       *http.ServeMux
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	storage      storage.Storage
	log          *slog.Logger
}

// NewServer create a new server
func NewServer(config *config.Config, store storage.Storage, log *slog.Logger) *server {
	s := &server{
		router:       http.NewServeMux(),
		Addr:         config.Address,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
		log:          log,
		storage:      store,
	}
	s.configureRouter()
	return s
}

// ServeHTTP кoutes HTTP requests using router
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter сonfigures server routing for commands.
func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleHome())
	s.router.HandleFunc(PathSave, s.handleSaveRunCommand(*s.log))
	s.router.HandleFunc(PathFind, s.handleGetOneCommand(*s.log))
	s.router.HandleFunc(PathList, s.handleGetListCommands(*s.log))
	s.router.HandleFunc(PathDelete, s.handleDeleteCommand(*s.log))
}

// error generates a error to the client
func (s *server) error(w http.ResponseWriter, code int, err error) {
	s.respond(w, code, map[string]string{"error": err.Error()})
}

// response generates a response to the client
func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
