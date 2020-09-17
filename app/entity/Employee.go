package entity
type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Fk_role    int    `json:"fk_role"`
}

type Authorization struct {
	Id    int
	Token string
	Login string
	Password string
	Role  string
}

