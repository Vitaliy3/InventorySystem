package mappers

import (
	"fmt"
	"log"
)

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Fk_role    int    `json:"fk_role"`
}

func (e *Employee) GetEmployeeById(id int) (employee Employee, err error) {
	OpenConnection()
	defer db.Close()
	fmt.Println("getId", id)
	row := db.QueryRow("select id,username,surname,patronymic,login,fk_role from users where id=$1", id)
	err = row.Scan(&e.Id, &employee.Name, &employee.Surname, &employee.Patronymic, &employee.Login, &employee.Fk_role)
	if err != nil {
		fmt.Println("errGetById:", err)
		return
	}
	return
}

func (e *Employee) GetAllEmployees() (employee []Employee, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select id,username,surname,patronymic,login,fk_role from users")
	if err != nil {
		fmt.Println("errGetById:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.Id, &e.Name, &e.Surname, &e.Patronymic, &e.Login, &e.Fk_role)
		if err != nil {
			log.Println(err)
			continue
		}
		employee = append(employee, *e)
	}
	return
}
func (e *Employee) UpdateEmployee() (rowsAffected int64, err error) {
	OpenConnection()
	defer db.Close()
	result, err := db.Exec("update users set username=$1 ,surname=$2 ,patronymic=$3,login=$4 where id=$5", e.Name, e.Surname, e.Patronymic, e.Login, e.Id)
	if err != nil {
		fmt.Println("err in updateEmployeeMapper", err)
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}
func (e *Employee) ResetPassEmployee(id int) (rowsAffected int64, err error) {
	OpenConnection()
	defer db.Close()
	fmt.Println("resetPasswordId", id)
	result, err := db.Exec("update users set userPassword=123 where id=$1", id)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

func (e *Employee) AddEmployee() (lastInsertedId int, err error) {
	OpenConnection()
	defer db.Close()
	err = db.QueryRow("insert into users (username,surname,patronymic,login,userpassword,fk_role) values($1,$2,$3,$4,$5,$6) returning id",
		e.Name, e.Surname, e.Patronymic, e.Login, e.Password, 2).Scan(&lastInsertedId)
	return
}
func (e *Employee) DeleteEmployee(id int) (rowsAffected int64, err error) {
	OpenConnection()
	defer db.Close()
	fmt.Println("getId", id)
	result, err := db.Exec("delete from users where id=$1", id)
	if err != nil {
		fmt.Println("errGetById:", err)
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}
