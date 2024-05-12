package storage

import "errors" // Перенести в сторадж // изменить нейминг

var ( //TODO: переписать пож себя названия, и нажо ли мен вообще они
	ErrCommandExists   = errors.New("command already exists")
	ErrCommandNotFound = errors.New("command not found")
	ErrMethod          = errors.New("invalid request method") //!!!скорее всего не сюда, переместить, поправить в хендлерах
	ErrEmptyRequest    = errors.New("empty request")
)
