package server

import (
	"encoding/json"
	"fmt"
	"main/internal/config"
	"main/internal/model"
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
	s.router.HandleFunc("/test", s.handleHome())

	http.HandleFunc("/", s.handleHome())
	http.HandleFunc("command/", handleSaveRunScript(s, "fjjgfkgjfkj")) // !!! Возможно надо прокинуть лог как у тузова
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }

	// s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	// s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
}

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
		// io.WriteString(w, "Hello World")
	}
}

type ScriptSave interface { //временно тут !!!
	SaveRunScript(name, script string) (int64, error)
}

func handleSaveRunScript( /*log *slog.Logger,*/ s *server, script ScriptSave) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const op = "handlers.Url.save.New"

		// log = log.With(
		// 	slog.String("op", op),
		// 	// slog.String("request_id", middleware.RequestID(r)), // r.Context
		// )
		var req *model.Request
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// s.error(w, r, http.StatusBadRequest, err)
			fmt.Println("Ошибка   op", op) // прекрутить проектную error.
			/// нужно отправить ответ с соответствующем кодом ошибки
			// log.Error("failed to decode request", sl.Err(err)) //если по тузову делать или
			// s.error(w, r, http.StatusBadRequest, err) // ошибка возможно гариллос выбрасывает

			return
		}
		// log.Info("request body decoded", slog.Any("request", req)) //Тузов

		// у основного создаеться здесь еще одна структура с данными юзер и присваиваев значения с запроса
		// u := &model.User{
		// 	Email:    req.Email,
		// 	Password: req.Password,
		// }
		if err := script.SaveRunScript(req.Name, req.Script); err != nil {

		}
		// if err := s.store.User().Create(u); err != nil {
		// 	s.error(w, r, http.StatusUnprocessableEntity, err)
		// 	return
		// }
		// s.respond(w, r, http.StatusCreated, u) // она реализована у осн

	}
}
