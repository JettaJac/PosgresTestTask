package server

import (
	"encoding/json"
	"errors"
	"main/internal/model"
	"main/internal/scripts"
	"main/internal/storage"
	"net/http"
	// "strconv"
	"io"
	"log/slog"
	"main/internal/lib/logger"
)

// import "net/http"
// curl -X POST -H "Content-Type: application/json" -d '{"name":"test555","script":"#!/bin/bash\necho \"Hello, World\""}' http://127.0.0.1:8080/command
// curl -X GET "Content-Type: application/json" -d '{"id":66}' http://127.0.0.1:8080/command

// type Command interface { //!!возможно не надо совсем
// 	SaveRunScript(urlTOSave, alias string) (int, error)
// 	GetOneScript(id int) (string, error)
// }

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(205)
		s.respond(w, r, 207, "req")
		w.Write([]byte("Hello NewTestHendler"))
		// io.WriteString(w, "Hello World")

	}
}

func (s *server) handleSaveRunCommand( /*log *slog.Logger, s *server*/ ) http.HandlerFunc { // TODO: возможно сделать как метод сервер в начале в скобках
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			const op = "server.handleSaveRunCommand"

			s.log = s.log.With(
				slog.String("op", op),
				// slog.String("request_id", middleware.RequestID(r)), // r.Context
			)
			var req *model.Command

			err := json.NewDecoder(r.Body).Decode(&req)
			if errors.Is(err, io.EOF) {
				// Такую ошибку встретим, если получили запрос с пустым телом.
				// Обработаем её отдельно
				s.log.Error("request body is empty")
				s.error(w, r, http.StatusUnprocessableEntity, err) // "empty request", возможнно добать в сообщение
			}

			if err != nil {
				s.log.Error("failed to decode request", sl.Err(err))
				s.error(w, r, http.StatusBadRequest, err)
			}

			s.log.Info("request body decoded", slog.Any("request", req))

			// у основного создаеться здесь еще одна структура с данными юзер и присваиваев значения с запроса
			// u := &model.User{
			// 	Email:    req.Email,
			// 	Password: req.Password,
			// }
			resScript, err := scripts.Run(req.Script) /// TODO:  Написать тесты на эту функцию, по типу тестов бево и валидация
			if err != nil {
				s.log.Error("failed to run command", sl.Err(err))
				s.error(w, r, http.StatusUnprocessableEntity, err)

			}
			req.Result = string(resScript) // !!! убрать стринг

			id, err := s.storage.SaveRunScript(req) //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд
			if errors.Is(err, storage.ErrURLExists) {
				s.log.Info("command already exists", slog.String("command", req.Script))
				s.error(w, r, http.StatusConflict, err)

			}

			if err != nil {
				s.log.Error("failed to add command", sl.Err(err))
				s.error(w, r, http.StatusUnprocessableEntity, err)
			}

			s.log.Info("command added", slog.Int("id", id)) //TODO: прикрутить лог тузова и раскомментировать
			s.respond(w, r, http.StatusCreated, req)        // она реализована у осн

		} else {
			s.log.Error("incorrect request method, need a POST")
			s.error(w, r, http.StatusMethodNotAllowed, storage.ErrMethod)
		}
	}
}

func (s *server) handleGetOneCommand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "server.handleGetOneCommand"
			s.log = s.log.With(
				slog.String("op", op),
				// slog.String("request_id", middleware.RequestID(r)), // r.Context
			)

			var req *model.Command
			var id int ///Запрос по id сделать

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				s.log.Error("failed to decode request")
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			s.log.Info("request body decoded", slog.Any("request", req))

			// idStr := r.URL.Query().Get("id") // Получаем значение параметра id из URL
			// 	if idStr != "" {
			// 				fmt.Fprintf(w, "Значение параметра id: %s", idStr )// !!!
			// 				return
			// 	} else {
			// 		fmt.Fprintf(w, "Параметр id %s не найден", idStr) // !!!
			// 		return
			// 	}
			// 	fmt.Println(idStr)

			err := s.storage.GetOneScript(req)

			if errors.Is(err, storage.ErrURLNotFound) {
				s.log.Info("command not found", slog.String("command", req.Script))
				s.error(w, r, http.StatusNotFound, err)
			}

			if err != nil {
				s.log.Error("failed to get command by id", sl.Err(err))
				s.error(w, r, http.StatusUnprocessableEntity, err) //!!! or http.StatusInternalServerError
			}

			s.log.Info("got command", slog.Int("id", id))
			s.respond(w, r, http.StatusOK, req)

		} else {
			s.log.Error("incorrect request method, need a GET")
			s.error(w, r, http.StatusMethodNotAllowed, storage.ErrMethod)
		}
	}
}

func (s *server) handleGetListCommands() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "server.handleGetListCommands"
			s.log = s.log.With(
				slog.String("op", op),
				// slog.String("request_id", middleware.RequestID(r)), // r.Context
			)

			listCommands, err := s.storage.GetListCommands() //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд
			if err != nil {
				s.log.Error("failed to add command", sl.Err(err))
				s.error(w, r, http.StatusInternalServerError, err)
			}

			s.log.Info("command list uploaded")
			s.respond(w, r, http.StatusOK, listCommands)
		} else {
			s.log.Error("incorrect request method, need a GET")
			s.error(w, r, http.StatusMethodNotAllowed, storage.ErrMethod)
		}
	}
}

func (s *server) handleDeleteCommand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			const op = "server.handleDeleteCommand"
			s.log = s.log.With(
				slog.String("op", op),
				// slog.String("request_id", middleware.RequestID(r)), // r.Context
			)

			// idStr := r.URL.Query().Get("id") // Получаем значение параметра id из URL
			// if idStr != "" {
			// 			fmt.Fprintf(w, "Значение параметра id: %s", idStr )// !!!
			// 			return
			// } else {
			// 	fmt.Fprintf(w, "Параметр id %s не найден", idStr) // !!!
			// 	return
			// }

			type Command struct {
				ID     int    `json:"id"`
				Result string `json:"result"`
			}
			// !!! здесь можно расширить если параметр id или скрипт, по нему и удалить

			var req Command //  возможно сделать в каждом хендере свою структуру запроса
			// !!! http.StatusNotFound
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				s.log.Error("failed to decode request")
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			s.log.Info("request body decoded", slog.Any("request", req))

			// id, _ := strconv.Atoi(idStr) //!!! обработать ошибку
			id := req.ID                       //Пока так, попробовать все таки с браузера забирать
			err := s.storage.DeleteCommand(id) //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд

			if errors.Is(err, storage.ErrURLNotFound) {
				s.log.Info("command not found", slog.String("command", "req.ID")) /// !!! преобразовать id в стринг
				s.error(w, r, http.StatusNotFound, err)
			}

			if err != nil {
				s.log.Error("failed to delete command by id", sl.Err(err))
				s.error(w, r, http.StatusInternalServerError, err) //!!! or http.StatusInternalServerError
			}

			req.ID = id
			req.Result = "No data, script deleted"
			s.log.Info("script deleted", slog.Int("id", id))
			s.respond(w, r, http.StatusOK, req) // она реализована у осн
		} else {
			s.log.Error("incorrect request method, need a ВУДУЕУ")
			s.error(w, r, http.StatusMethodNotAllowed, storage.ErrMethod)
		}
	}
}
