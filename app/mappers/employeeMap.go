package mappers

import (
	"database/sql"
	"myapp/app/entity"
)

type Employee struct {
	entity.Employee
}

//получение роли пользователя
func (e *Employee) GetUserRoleById(DB *sql.DB, id int) (role string, err error) {
	row := DB.QueryRow("select r.userRole from users u join roles r on u.fk_role =r.id where u.id =$1", id)
	err = row.Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//получение данных о пользователе
func (e *Employee) Auth(DB *sql.DB, login string) (employee Employee, err error) {
	row := DB.QueryRow("select id,login,userpassword,fk_role from users  where login =$1;", login)
	err = row.Scan(&employee.Id, &employee.Login, &employee.Password, &employee.Fk_role)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//получение данных о сотруднике
func (e *Employee) GetEmployeeById(DB *sql.DB, id int) (employee entity.Employee, err error) {
	row := DB.QueryRow("select id,username,surname,patronymic,login,fk_role from users where id=$1", id)
	err = row.Scan(&employee.Id, &employee.Name, &employee.Surname, &employee.Patronymic, &employee.Login, &employee.Fk_role)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//выборка всех сотрудников
func (e *Employee) GetAllEmployees(DB *sql.DB) (employees []entity.Employee, err error) {
	rows, err := DB.Query("select id,username,surname,patronymic,login,fk_role from users where fk_role=2")
	if err != nil {
		return
	}
	var employee entity.Employee
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Surname, &employee.Patronymic, &employee.Login, &employee.Fk_role)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			return
		}
		employees = append(employees, employee)
	}
	return
}

func (e *Employee) UpdateEmployee(DB *sql.DB,employee entity.Employee) (lastUpdateId int, err error) {
	err = DB.QueryRow("update users set username=$1 ,surname=$2 ,patronymic=$3,login=$4 where id=$5 returning id", employee.Name, employee.Surname, employee.Patronymic, employee.Login, employee.Id).Scan(&lastUpdateId)
	if err != nil {
		return
	}
	return
}

//сброс пароля сотрудника
func (e *Employee) ResetPassEmployee(DB *sql.DB, employee entity.Employee) (updatedRowId int, err error) {
	err = DB.QueryRow("update users set userPassword=123 where id=$1 returning id", employee.Id).Scan(&updatedRowId)
	if err != nil {
		return
	}
	return
}

//добавление сотрудника
func (e *Employee) AddEmployee(DB *sql.DB,employee entity.Employee) (lastInsertedId int, err error) {
	err = DB.QueryRow("insert into users (username,surname,patronymic,login,userpassword,fk_role) values($1,$2,$3,$4,$5,$6) returning id",
		employee.Name, employee.Surname, employee.Patronymic, employee.Login, employee.Password, 2).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

//удаление сотрудника
func (e *Employee) DeleteEmployee(DB *sql.DB, employee entity.Employee) (lastDeleteId int, err error) {
	err = DB.QueryRow("delete from users where id=$1 returning id", employee.Id).Scan(&lastDeleteId)
	if err != nil {
		return
	}
	return
}
