package storage

import "main/internal/model"

type Storage interface { /// TODO: возможно переименновать в CommandRun
	SaveRunScript(req *model.Command) (int64, error)
	GetOneScript(req *model.Command) (string, error)
	GetListScrits(req *model.Command) ([]string, error)
}