package handlers

import "net/http"

type Commands interface {
	SaveRunScript(urlTOSave, alias string) (int64, error)
}

func HandleSaveRunScript() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
