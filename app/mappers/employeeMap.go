package mappers

import "database/sql"

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Fk_role    int    `json:"fk_role"`
}

func (e *Employee) GetEmployeeById(DB *sql.DB,id int) (employee Employee, err error) {
	row := DB.QueryRow("select id,username,surname,patronymic,login,fk_role from users where id=$1", id)
	err = row.Scan(&e.Id, &employee.Name, &employee.Surname, &employee.Patronymic, &employee.Login, &employee.Fk_role)
	if err != nil{
		return
	}
	return
}

func (e *Employee) GetAllEmployees(DB *sql.DB) (employee []Employee, err error) {
	rows, err := DB.Query("select id,username,surname,patronymic,login,fk_role from users")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Id, &e.Name, &e.Surname, &e.Patronymic, &e.Login, &e.Fk_role)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			return
		}
		employee = append(employee, *e)
	}
	return
}

func (e *Employee) UpdateEmployee(DB *sql.DB) (lastUpdateId int, err error) {
	err = DB.QueryRow("update users set username=$1 ,surname=$2 ,patronymic=$3,login=$4 where id=$5 returning id", e.Name, e.Surname, e.Patronymic, e.Login, e.Id).Scan(&lastUpdateId)
	if err != nil {
		return
	}
	return
}
func (e *Employee) ResetPassEmployee(DB *sql.DB, id int) (updatedRowId int, err error) {
	err = DB.QueryRow("update users set userPassword=123 where id=$1 returning id", id).Scan(&updatedRowId)
	if err != nil {
		return
	}
	return
}

func (e *Employee) AddEmployee(DB *sql.DB) (lastInsertedId int, err error) {
	err = DB.QueryRow("insert into users (username,surname,patronymic,login,userpassword,fk_role) values($1,$2,$3,$4,$5,$6) returning id",
		e.Name, e.Surname, e.Patronymic, e.Login, e.Password, 2).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}
func (e *Employee) DeleteEmployee(DB *sql.DB, id int) (lastDeleteId int, err error) {
	err = DB.QueryRow("delete from users where id=$1 returning id", id).Scan(&lastDeleteId)
	if err != nil {
		return
	}
	return
}
