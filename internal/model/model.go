package model

// Command ... //TODO: возможно стелать как в осн , черезе репозиторий интерфейс В2-15.30
type Command struct {
	ID     int    `json:"id" db:"id"`
	Script string `json:"script" db:"script"` // возможно прикрутить валидацию
	Result string `json:"result" db:"result"`
	// TODO:   если успею прикрутить авторизацию /юзера и пароль
}

// type Request struct {
// 	// ID       int    json:"id"
// 	Name   string `json:"name"`
// 	Script string `json:"script"` // возможно прикрутить валидацию
// 	// TODO:   если успею прикрутить авторизацию /юзера и пароль
// }

// type Response struct { //TODO:  вроде не используеться
// 	ID     int    `json:"id"`
// 	Status string `json:"status"`
// 	Script string `json:"script,omitempty"`
// 	// Error  string `json:"error,omitempty"` //скорее всего стату решает
// 	Result string `json:"result,omitempty"` // возможноне нежно выводить, подумать
// }
