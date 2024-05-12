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

// import "net/http"
// curl -X POST -H "Content-Type: application/json" -d '{"name":"test555","script":"#!/bin/bash\necho \"Hello, World\""}' http://127.0.0.1:8080/command
// curl -X GET "Content-Type: application/json" -d '{"id":66}' http://127.0.0.1:8080/command

func (s *server) handleSaveRunCommand(log slog.Logger /*, s *server*/) http.HandlerFunc { // TODO: возможно сделать как метод сервер в начале в скобках
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
				s.error(w, http.StatusBadRequest, err) // "empty request", возможнно добать в сообщение
				return
			}

			if err != nil {
				s.log.Error("failed to decode request", sl.Err(err))
				s.error(w, http.StatusBadRequest, err)
				return
			}
			err = model.ValidateJson(req)
			if err != nil {
				// Обработка ошибки валидации
				// err.(validator.ValidationErrors) содержит информацию о полях, которые не прошли валидацию
				// Можно обработать ошибку и вывести информацию о невалидных полях
				s.log.Error("incorrect JSON request", sl.Err(err))
				s.error(w, http.StatusBadRequest, err)
				return
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
				s.error(w, http.StatusUnprocessableEntity, err)
				return

			}
			s.log.Info("command run ")
			req.Result = string(resScript) // !!! убрать стринг

			id, err := s.storage.SaveRunScript(req) //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд
			// fmt.Println(err)
			// if errors.Is(err, fmt.Errorf("%s: %s", "storage.SaveRunScript", storage.ErrCommandExists)) {
			// 	s.log.Error("command already exists", slog.String("command", req.Script))
			// 	s.error(w, http.StatusConflict, err)
			// 	return
			// }

			if err != nil {
				s.log.Error("failed to add command", sl.Err(err))
				s.error(w, http.StatusBadRequest, err)
				return
			}

			s.log.Info("command added", slog.Int("id", id)) //TODO: прикрутить лог тузова и раскомментировать
			s.respond(w, http.StatusCreated, req)           // она реализована у осн

		} else {
			s.log.Error("incorrect request method, need a POST")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}

func (s *server) handleGetOneCommand(log slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "server.handleGetOneCommand"
			log = *s.log.With(
				slog.String("op", op),
			)

			var id int

			///Запрос по id сделать
			idStr := r.URL.Query().Get("id") // Получаем значение параметра id из URL

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

			req, err := s.storage.GetOneScript(id)
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

func (s *server) handleGetListCommands(log slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "server.handleGetListCommands"
			log = *s.log.With(
				slog.String("op", op),
			)

			listCommands, err := s.storage.GetListCommands() //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд
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

			idStr := r.URL.Query().Get("id") // Получаем значение параметра id из URL
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
			// !!! здесь можно расширить если параметр id или скрипт, по нему и удалить

			err := s.storage.DeleteCommand(req.ID) //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд

			if errors.Is(err, storage.ErrCommandNotFound) {
				s.log.Error("command not found", slog.String("command", strconv.Itoa(req.ID))) /// !!! преобразовать id в стринг
				s.error(w, http.StatusNotFound, err)
				return
			}

			if err != nil {
				s.log.Error("failed to delete command by id", sl.Err(err))
				s.error(w, http.StatusInternalServerError, err) //!!! or http.StatusInternalServerError
				return
			}

			req.Result = "script deleted"
			s.log.Info("script deleted", slog.Int("id", req.ID))
			s.respond(w, http.StatusOK, req) // она реализована у осн
		} else {
			s.log.Error("incorrect request method")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// listCommands, err := s.storage.GetListCommands()
			// _, _ = listCommands, err
			s.respond(w, 200, "Начнем!!!")
			w.Write([]byte("Hello World!"))
		} else {
			s.log.Error("incorrect request method, need a POST")
			s.error(w, http.StatusMethodNotAllowed, storage.ErrMethod)
			return
		}
	}
}
