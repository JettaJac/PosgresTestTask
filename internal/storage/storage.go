package storage

import (
	"errors"
	"main/internal/model"
)

var (
	ErrCommandExists   = errors.New("command already exists")
	ErrCommandNotFound = errors.New("command not found")
	ErrMethod          = errors.New("invalid request method")
	ErrEmptyRequest    = errors.New("empty request")
)

// Storage ...
type Storage interface { /// TODO: возможно переименновать в CommandRun
	SaveRunCommand(req *model.Command) (int, error)
	GetOneCommand(id int) (*model.Command, error) // !!!исправить везде нейминг
	GetListCommands() ([]model.Command, error)
	DeleteCommand(id int) error
	// TODO: реализовать методы поиск и  удаление по скрипту
}
