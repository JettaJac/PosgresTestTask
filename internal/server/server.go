package server

import (
	"encoding/json"
	"main/internal/config"
	"main/internal/storage/sqlstore"
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
	storage      *sqlstore.Storage
	// log       *slog.Logger
	// storage storage.Storage
}

func NewServer(config *config.Config, store *sqlstore.Storage /*, log *slog.Logger,sessionStore sessions.Store*/) /**server*/ {
	s := &server{
		router:       http.NewServeMux(),
		Addr:         config.Address,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
		// log: log // logger: logrus.New(),
		storage: store,
		// sessionStore: sessionStore,
	}
	s.configureRouter()
	// return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/test", s.handleHome())

	http.HandleFunc("/", s.handleHome())
	// http.HandleFunc("command/", h.HandleSaveRunScript(s)) // !!! Возможно надо прокинуть лог как у тузова
	http.HandleFunc("command/", handleSaveRunScript( /*log, */ s)) // !!! Возможно надо прокинуть лог как у тузова// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }

	// s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
}

type ScriptSave interface { //временно тут !!! //!!возможно не надо совсем
	SaveRunScript(name, script string) (int64, error)
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
