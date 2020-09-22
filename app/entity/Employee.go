package entity

//структура сотрудника
type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`       //имя
	Surname    string `json:"surname"`    //фамилия
	Patronymic string `json:"patronymic"` //отчество
	Login      string `json:"login"`      //логин
	Password   string `json:"password"`   //пароль
	Fk_role    int    `json:"fk_role"`    //внешний ключ на роль
}

//структура для авторизации
type Authorization struct {
	Id       int
	Token    string //токен
	Login    string //логин
	Password string //пароль
	Role     string //роль
}
