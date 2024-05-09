package save

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	sl "main/internal/lib/logger"
	"main/internal/storage"
	"net/http"
)

// !!!вынести отдельно в модель-?
type Request struct {
	// ID       int    json:"id"
	Name   string `json:"name"`
	Script string `json:"script"` // возможно прикрутить валидацию
	// Result   string `json:"result,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}

type ScriptSave interface {
	SaveScript(urlTOSave, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver ScriptSave) http.HandlerFunc { // hendler
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Url.save.New"

		log = log.With(
			slog.String("op", op),
			// slog.String("request_id", middleware.RequestID(r)), // r.Context
		)
		var req *Request
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			// s.error(w, r, http.StatusBadRequest, err)
			fmt.Println("Ошибка   op") // прекрутить проектную error.
			/// нужно отправить ответ с соответствующем кодом ошибки
			log.Error("failed to decode request", sl.Err(err))

			return
		}
		// err :=

		// err:= render.DecodeJson(r.Body, &req)
		// if err != nil{
		log.Info("request body decoded", slog.Any("request", req))

		// if err :=validator.New().Struct(req); err!= nil {
		// 	log.Error("failed to validate request", sl.Err(err))
		// 	//
		// 	return

		// }

		id, err := urlSaver.SaveScript(req.Name, req.Script)
		if errors.Is(err, storage.ErrCommandExists) {
			log.Info("url already exists", slog.String("url", req.Name))
			// генерирует responced
			return
		}
		if err != nil {
			log.Error("failed to add url", sl.Err(err))
			// генерирует responced  у Тузова  render.JSON*w,r,resp.Error("failed to add url")
			return
		}

		log.Info("url added", slog.Int64("id", id))
		// response(w,r,alias)// тузов

	}

}

// func responseOk(w http.ResponseWriter, r *http.Request, alias string) {// Тузов
// 	render.JSON(w, r, Response{
// 		Response: resp.Ok(),
// 		Alias: alias,
// 	})
// }
