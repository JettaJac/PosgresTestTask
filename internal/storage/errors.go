package storage

import "errors" // Перенести в сторадж // изменить нейминг

var ( //TODO: переписать пож себя названия, и нажо ли мен вообще они
	ErrCommandExists   = errors.New("Command already exists")
	ErrCommandNotFound = errors.New("Command not found")
	ErrMethod          = errors.New("Invalid request method") //!!!скорее всего не сюда, переместить, поправить в хендлерах
	ErrEmptyRequest    = errors.New("Empty request")
)

// type Command struct {
//  ID       int    json:"id"
//  Name     string json:"name"
//  Script   string json:"script"
//  Result   string json:"result"
// }
