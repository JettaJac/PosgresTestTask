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

type server struct {
	router       *http.ServeMux
	Addr         string // теоритиически можно убрать
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	storage      storage.Storage
	log          *slog.Logger
	// storage storage.Storage
}

func NewServer(config *config.Config, store storage.Storage, log *slog.Logger /*,sessionStore sessions.Store*/) *server {
	s := &server{
		router:       http.NewServeMux(),
		Addr:         config.Address,
		ReadTimeout:  config.HTTPServer.Timeout,
		WriteTimeout: config.HTTPServer.Timeout,
		IdleTimeout:  config.HTTPServer.IdleTimeout,
		log:          log, // logger: logrus.New(),
		// storage: store,
		storage: store,
		// sessionStore: sessionStore,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/test", s.handleHome())

	http.HandleFunc("/", s.handleHome())
	// .Methods("POST")
	// http.HandleFunc("command/", h.HandleSaveRunScript(s)) // TODO:  Возможно надо прокинуть лог как у тузова
	http.HandleFunc("/command/save", s.handleSaveRunCommand( /*log, */ )) // TODO:  Возможно надо прокинуть лог как у тузова// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }
	http.HandleFunc("/command/find", s.handleGetOneCommand())
	http.HandleFunc("/commands/all", s.handleGetListCommands())
	http.HandleFunc("/command/delete", s.handleDeleteCommand())

	// s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
}

// type ScriptSave interface { //временно тут TODO:  //!!возможно не надо совсем
// 	SaveRunScript(name, script string) (int64, error)
// }

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
	return
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	// w.Write([]byte("Hello TestHendler"))
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
