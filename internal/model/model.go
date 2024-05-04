package model

// Command ... //!!!возможно стелать как в осн , черезе репозиторий интерфейс В2-15.30
type Command struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Script string `json:"script"` // возможно прикрутить валидацию
	Result string `json:"result"`
	// !!!  если успею прикрутить авторизацию /юзера и пароль
}

type Request struct {
	// ID       int    json:"id"
	Name   string `json:"name"`
	Script string `json:"script"` // возможно прикрутить валидацию
	// !!!  если успею прикрутить авторизацию /юзера и пароль
}

type Response struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Script string `json:"script,omitempty"`
	// Error  string `json:"error,omitempty"` //скорее всего стату решает
	Result string `json:"result,omitempty"` // возможноне нежно выводить, подумать
}
