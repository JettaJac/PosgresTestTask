package server

import (
	"encoding/json"
	"main/internal/config"
	// "main/internal/storage/sqlstore"
	"main/internal/storage"
	"net/http"
	"time"
	// "main/internal/config"
	"log/slog"
)

var (
	PathSave   = "/command/save"   // handleSaveRunCommand
	PathFind   = "/command/find"   //HandleGetOneCommand
	PathList   = "/commands/all"   // HandleGetListCommands
	PathDelete = "/command/delete" // HandleDeleteCommand
)

type server struct {
	router       *http.ServeMux
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	storage      storage.Storage
	log          *slog.Logger
}

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

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
} /// скорее всего не нужно

func (s *server) configureRouter() {
	s.router.HandleFunc("/", s.handleHome())
	s.router.HandleFunc(PathSave, s.handleSaveRunCommand(*s.log)) // TODO:  Возможно надо прокинуть лог как у тузова// err := http.ListenAndServe(":8000", nil)
	s.router.HandleFunc(PathFind, s.handleGetOneCommand(*s.log))
	s.router.HandleFunc(PathList, s.handleGetListCommands(*s.log))
	s.router.HandleFunc(PathDelete, s.handleDeleteCommand(*s.log))
}

func (s *server) error(w http.ResponseWriter /*r *http.Request, */, code int, err error) {
	s.respond(w, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
