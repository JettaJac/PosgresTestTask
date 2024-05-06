package storage

import (
	"main/internal/model"
"net/http"
)


type Storage interface { /// TODO: возможно переименновать в CommandRun
	SaveRunScript(req *model.Command) (int, error)
	GetOneScript(req *model.Command) (/*string,*/ error)
	GetListCommands(req *model.Command,  w http.ResponseWriter) ([]model.Command, error)
	DeleteCommand(id int) error
}