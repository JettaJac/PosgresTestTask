package storage

import (
	"main/internal/model"
)

type Storage interface { /// TODO: возможно переименновать в CommandRun
	SaveRunScript(req *model.Command) (int, error)
	GetOneScript(id int) (*model.Command, error) // !!!исправить везде нейминг
	GetListCommands() ([]model.Command, error)
	DeleteCommand(id int) error
	// TODO: реализовать мктоды поиск и  удаление по скрипту
}
