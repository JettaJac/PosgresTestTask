package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"main/internal/model"
	"main/internal/scripts"
	"main/internal/storage"
	"net/http"
)

// import "net/http"
// curl -X POST -H "Content-Type: application/json" -d '{"name":"test555","script":"#!/bin/bash\necho \"Hello, World\""}' http://127.0.0.1:8080/command
// curl -X GET "Content-Type: application/json" -d '{"id":66}' http://127.0.0.1:8080/commands

// type Command interface { //!!возможно не надо совсем
// 	SaveRunScript(urlTOSave, alias string) (int64, error)
// 	GetOneScript(id int64) (string, error)
// }

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello TestHendler"))
		// io.WriteString(w, "Hello World")
	}
}

func (s *server) handleSaveRunCommand( /*log *slog.Logger, s *server*/) http.HandlerFunc { // TODO: возможно сделать как метод сервер в начале в скобках
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO:  сделать, чтоб  Postзапрос обрабатывал
		fmt.Println(r.Method, "save test", http.MethodPost )
		if r.Method == http.MethodPost {

		// }
		fmt.Println("save test")
		const op = "handleSaveRunCommand"

		// log = log.With(
		// 	slog.String("op", op),
		// 	// slog.String("request_id", middleware.RequestID(r)), // r.Context
		// )
		var req *model.Command
		// fmt.Println(r)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			// fmt.Println(r)
			// fmt.Println("REQ: ", req)
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
		resScript, err := scripts.Run(req.Script) /// TODO:  Написать тесты на эту функцию, по типу тестов бево и валидация
		if err != nil {
			// log.Error("failed to add url", sl.Err(err)) TODO: прикрутить лог тузова и раскомментировать

			s.error(w, r, http.StatusUnprocessableEntity, err)
			// генерирует responced  у Тузова  render.JSON*w,r,resp.Error("failed to add url")
			return
		}
		req.Result = string(resScript)

		id, err := s.storage.SaveRunScript(req) //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд
		fmt.Println(id)                         //  TODO:  tmp
		if errors.Is(err, storage.ErrURLExists) {
			// log.Info("url already exists", slog.String("url", req.Name)) TODO: прикрутить лог тузова и раскомментировать
			// генерирует responced
			s.error(w, r, http.StatusConflict, err) //  TODO:  проверить правильность статуса (Tuz)
			return
		}

		if err != nil {
			// log.Error("failed to add url", sl.Err(err)) TODO: прикрутить лог тузова и раскомментировать

			s.error(w, r, http.StatusUnprocessableEntity, err)
			// генерирует responced  у Тузова  render.JSON*w,r,resp.Error("failed to add url")
			return
		}
		req.Result = string(resScript) //TODO:  возможно изменить тип, чтоб не делать преобразований скорее всего так нельзя, надо сделать запрос в бд и посмотреть записалось ли туда
		//// if err := s.store.User().Create(u); err != nil {
		// 	s.error(w, r, http.StatusUnprocessableEntity, err)
		// 	return
		// }
		// log.Info("url added", slog.Int64("id", id)) TODO: прикрутить лог тузова и раскомментировать
		s.respond(w, r, http.StatusCreated, req) // она реализована у осн

	} else {
		// выкинуть ошибку о не правильном методе запроса что то типо  http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		//Применить во всех запросах
	}
}
}

func (s *server) handleGetOneCommand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "handleGetOneCommand"

		var req *model.Command

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println("Ошибка   op", op) // прекрутить проектную error.
			return
		}


		_, err := s.storage.GetOneScript(req) //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд
		// fmt.Println(id)                         //  TODO:  tmp
		if errors.Is(err, storage.ErrURLExists) { //  TODO: посмотреть, возможно нужно также отдельно проверить, но другуюю ошибку, так как тут ищем, а не создаем
			// log.Info("url already exists", slog.String("url", req.Name)) TODO: прикрутить лог тузова и раскомментировать
			// генерирует responced
			s.error(w, r, http.StatusConflict, err) //  TODO:  проверить правильность статуса (Tuz)
			return
		}

		if err != nil {
			// log.Error("failed to add url", sl.Err(err)) TODO: прикрутить лог тузова и раскомментировать

			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		// log.Info("url added", slog.Int64("id", id)) TODO: прикрутить лог тузова и раскомментировать
		s.respond(w, r, http.StatusCreated, req) // она реализована у осн

// TODO:логи привести к одному виду, одинаковая структура логов в функциях
		}
		
	}
}

func (s *server) handleGetListCommands() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			const op = "handleGetListCommands"

		var req *model.Command //  возможно сделать в каждом хендере свою структуру запроса

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println("Ошибка   op", op) // прекрутить проектную error.
			return
		}


		_, err := s.storage.GetOneScript(req) //TODO:  нужен интервейс, так не хорошо прокидывать напряму бд
		// fmt.Println(id)                         //  TODO:  tmp
		if errors.Is(err, storage.ErrURLExists) { //  TODO: посмотреть, возможно нужно также отдельно проверить, но другуюю ошибку, так как тут ищем, а не создаем
			// log.Info("url already exists", slog.String("url", req.Name)) TODO: прикрутить лог тузова и раскомментировать
			// генерирует responced
			s.error(w, r, http.StatusConflict, err) //  TODO:  проверить правильность статуса (Tuz)
			return
		}

		if err != nil {
			// log.Error("failed to add url", sl.Err(err)) TODO: прикрутить лог тузова и раскомментировать

			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		// log.Info("url added", slog.Int64("id", id)) TODO: прикрутить лог тузова и раскомментировать
		s.respond(w, r, http.StatusCreated, req) // она реализована у осн

// TODO:логи привести к одному виду, одинаковая структура логов в функциях
		}
		
	}
}