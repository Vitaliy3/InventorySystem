package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"myapp/app/mappers"
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

//получение всех сотрудников
func (e *Employee) GetAllEmployees(DB *sql.DB) (employeeArray []Employee, err error) {
	emMapper := mappers.Employee{}
	dbEmployees, err := emMapper.GetAllEmployees(DB)
	if err != nil {
		return
	}
	var temp Employee
	for _, v := range dbEmployees {
		temp.Id = v.Id
		temp.Name = v.Name
		temp.Surname = v.Surname
		temp.Patronymic = v.Patronymic
		temp.Login = v.Login
		employeeArray = append(employeeArray, temp)
	}
	return
}

//изменение данных о сотруднике
func (e *Employee) UpdateEmployee(DB *sql.DB,params *revel.Params) (employee Employee, err error) {
	employeeMapper := mappers.Employee{}
	err = json.Unmarshal(params.JSON, &employeeMapper)
	if err != nil {
		return
	}
	lastUpdateId, err := employeeMapper.UpdateEmployee(DB)
	if err != nil {
		return
	}
	employeeMapper, err = employeeMapper.GetEmployeeById(DB,lastUpdateId)
	employee.Id = employeeMapper.Id
	employee.Name = employeeMapper.Name
	employee.Surname = employeeMapper.Surname
	employee.Patronymic = employeeMapper.Patronymic
	employee.Login = employeeMapper.Login
	if err != nil {
		return
	}
	return
}

//удаление сотрудника
func (e *Employee) DeleteEmployee(DB *sql.DB,params *revel.Params) (employee Employee, err error) {
	emMapper := mappers.Employee{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}
	lastDeleteId, err := emMapper.DeleteEmployee(DB,id)
	if err != nil {
		return
	}
	employee.Id = lastDeleteId
	return
}

func (e *Employee) AddEmployee(DB *sql.DB,params *revel.Params) (employee Employee, err error) {
	employeeMapper := mappers.Employee{}
	err = json.Unmarshal(params.JSON, &employeeMapper)
	if err != nil {
		return
	}
	lastInsertedId, err := employeeMapper.AddEmployee(DB)
	if err != nil {
		return
	}
	employeeMapper, err = employeeMapper.GetEmployeeById(DB,lastInsertedId)
	if err != nil {
		return
	}
	employee.Id = employeeMapper.Id
	employee.Name = employeeMapper.Name
	employee.Surname = employeeMapper.Surname
	employee.Patronymic = employeeMapper.Patronymic
	employee.Login = employeeMapper.Login
	return
}

func (e *Employee) ResetPassEmployee(DB *sql.DB,params *revel.Params) (employee Employee, err error) {
	emMapper := mappers.Employee{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	updatedRowId, err := emMapper.ResetPassEmployee(DB,id)
	if err != nil {
		return
	}
	if id != updatedRowId {
		err = errors.New("not reset")
		return
	}
	return
}
