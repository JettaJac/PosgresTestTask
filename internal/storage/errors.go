package storage

import "errors"

var ( //TODO: переписать пож себя названия, и нажо ли мен вообще они
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExists   = errors.New("URL already exists")
	ErrMethod  = errors.New ("Invalid request method") //!!!скорее всего не сюда, переместить, поправить в хендлерах
)

// type Command struct {
//  ID       int    json:"id"
//  Name     string json:"name"
//  Script   string json:"script"
//  Result   string json:"result"
// }
