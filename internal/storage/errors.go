package storage

import "errors"

var ( //!!!переписать пож себя названия, и нажо ли мен вообще они
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExists   = errors.New("URL already exists")
)

// type Command struct {
//  ID       int    json:"id"
//  Name     string json:"name"
//  Script   string json:"script"
//  Result   string json:"result"
// }
