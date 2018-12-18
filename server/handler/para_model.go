package handler

type parameterLoginModel struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type parameterTodoModel struct {
	Id        int64  `json:"id""`
	Title     string `json:"title"`
	Completed string `json:"completed"`
}
