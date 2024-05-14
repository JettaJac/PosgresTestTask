package server

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"main/internal/lib/logger"
	"main/internal/model"
	"main/internal/scripts"
	"main/internal/storage"
	"net/http"
	"strconv"
)

// handleSaveRunCommand saves the command to the database
func (s *server) handleSaveRunCommand(log slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			const op = "server.handleSaveRunCommand"

			log = *s.log.With(
				slog.String("op", op),
			)
			var req *model.Command

			err := json.NewDecoder(r.Body).Decode(&req)
			if errors.Is(err, io.EOF) {
				s.log.Error("request body is empty")
				s.error(w, http.StatusBadRequest, err)
				return
			}

			if err != nil {
				s.log.Error("failed to decode request", sl.Err(err))
				s.error(w, http.StatusBadRequest, err)
				return
			}
			err = model.ValidateJson(req)
			if err != nil {
				s.log.Error("incorrect JSON request", sl.Err(err))
				s.error(w, http.StatusBadRequest, err)
				return
			}

			s.log.Info("request body decoded", slog.Any("request", req))

			resScript, err := scripts.Run(req.Script)
			if err != nil {
				s.log.Error("failed to run command", sl.Err(err))
				s.error(w, http.StatusUnprocessableEntity, err)
				return

			}
			s.log.Info("command run ")
			req.Result = string(resScript)

			id, err := s.storage.SaveRunCommand(req)
			if err != nil {
				s.log.Error("failed to add command", sl.Err(err))
				s.error(w, http.StatusBadRequest, err)
				return
			}

			s.log.Info("command added", slog.Int("id", id))
			s.respond(w, http.StatusCreated, req)

		} else {
			s.log.Error("incorrect request method, need a POST")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}

// / handleGetOneCommand returns one command from the database by ID
func (s *server) handleGetOneCommand(log slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "server.handleGetOneCommand"
			log = *s.log.With(
				slog.String("op", op),
			)

			var id int

			idStr := r.URL.Query().Get("id")

			if idStr != "" {
				var err error
				id, err = strconv.Atoi(idStr)
				if err != nil {
					s.log.Error("incorrect ID entered", slog.String("id: ", idStr))
					s.error(w, http.StatusBadRequest, err)
					return
				}

			} else {
				s.log.Error("incorrect ID entered", slog.String("id: ", idStr))
				s.error(w, http.StatusBadRequest, storage.ErrEmptyRequest)
				return
			}

			req, err := s.storage.GetOneCommand(id)
			if errors.Is(err, storage.ErrCommandNotFound) {
				s.log.Error("command not found", slog.String("command id: ", idStr))
				s.error(w, http.StatusNotFound, err)
				return
			}
			if err != nil {
				s.log.Error("failed to get command by id", sl.Err(err))
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.log.Info("got command", slog.Int("id", req.ID))
			s.respond(w, http.StatusOK, req)

		} else {
			s.log.Error("incorrect request method, need a GET")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}

// // handleGetListCommands returns all commands from the database as a list
func (s *server) handleGetListCommands(log slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "server.handleGetListCommands"
			log = *s.log.With(
				slog.String("op", op),
			)

			listCommands, err := s.storage.GetListCommands()
			if err != nil {
				s.log.Error("failed to add command", sl.Err(err))
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.log.Info("command list uploaded")
			s.respond(w, http.StatusOK, listCommands)
		} else {
			s.log.Error("incorrect request method, need a GET")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}

// / handleDeleteCommand deletes a command from the database by ID
// TODO: здесь можно расширить удалением по команде, по нему и удалить
func (s *server) handleDeleteCommand(log slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			const op = "server.handleDeleteCommand"
			log = *s.log.With(
				slog.String("op", op),
			)

			req := &model.Command{
				ID:     0,
				Script: "",
				Result: "",
			}

			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					s.log.Error("incorrect ID entered", slog.String("id: ", idStr))
					s.error(w, http.StatusBadRequest, err)
					return
				}
				req.ID = id
			} else {
				s.log.Error("incorrect ID entered", slog.String("id: ", idStr))
				s.error(w, http.StatusBadRequest, storage.ErrEmptyRequest)
				return
			}

			type Command struct {
				ID     int    `json:"id"`
				Result string `json:"result"`
			}

			err := s.storage.DeleteCommand(req.ID)
			if errors.Is(err, storage.ErrCommandNotFound) {
				s.log.Error("command not found", slog.String("command", idStr))
				s.error(w, http.StatusNotFound, err)
				return
			}

			if err != nil {
				s.log.Error("failed to delete command by id", sl.Err(err))
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			req.Result = "script deleted"
			s.log.Info("script deleted", slog.Int("id", req.ID))
			s.respond(w, http.StatusOK, req)
		} else {
			s.log.Error("incorrect request method")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}

// handleHome returns the home page
func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.respond(w, 200, "Начнем!!!")
			w.Write([]byte("Hello World!"))
		} else {
			s.log.Error("incorrect request method, need a POST")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}
