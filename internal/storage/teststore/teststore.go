package teststore

import (
	"main/internal/model"
	"main/internal/storage"
)

type Storage struct {
	Commands map[string]*model.Command
}

// New reate a new storage
func New() *Storage {
	return &Storage{
		Commands: make(map[string]*model.Command),
	}
}

// SaveRunCommand create a new command
func (s *Storage) SaveRunCommand(req *model.Command) (int, error) {
	if _, ok := s.Commands[req.Script]; !ok {
		s.Commands[req.Script] = req
		req.ID = len(s.Commands)

		return req.ID, nil
	}
	return 0, storage.ErrCommandExists
}

// / GetOneCommand get one command from storage by ID
func (s *Storage) GetOneCommand(id int) (*model.Command, error) {

	count := 0
	req := &model.Command{
		ID:     0,
		Script: "",
		Result: "",
	}

	for _, v := range s.Commands {
		if v.ID == id {
			req = v
			count++
			continue
		}
	}

	if count <= 0 {
		return nil, storage.ErrCommandNotFound
	}

	return req, nil
}

// GetListCommands get all commands from storage
func (s *Storage) GetListCommands() ([]model.Command, error) {
	var commands []model.Command

	for _, v := range s.Commands {
		commands = append(commands, *v)
	}

	return commands, nil
}

// DeleteCommand delete a command from storage by ID
func (s *Storage) DeleteCommand(id int) error {
	count := 0
	for _, v := range s.Commands {
		if v.ID == id {
			delete(s.Commands, v.Script)
			v = nil
			count++
		}
	}
	if count <= 0 {
		return storage.ErrCommandNotFound
	}

	return nil
}
