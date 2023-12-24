package entities

type Student struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Group     string `json:"group"`
	Course    string `json:"course"`
	Specialty string `json:"specialty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Result string `json:"result"`
}
