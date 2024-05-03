package model

type Request struct {
	// ID       int    json:"id"
	Name   string `json:"name"`
	Script string `json:"script"` // возможно прикрутить валидацию
	// !!!  если успею прикрутить авторизацию /юзера и пароль
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}
