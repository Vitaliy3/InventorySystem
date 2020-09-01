package models

import (
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
func (e *Employee) GetAllEmployees() (employeeArray []Employee, err error) {
	emMapper := mappers.Employee{}
	dbEmployees, err := emMapper.GetAllEmployees()
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
func (e *Employee) UpdateEmployee(params *revel.Params) (employee Employee, err error) {
	employeeMapper := mappers.Employee{}
	err = json.Unmarshal(params.JSON, &employeeMapper)
	if err != nil {
		return
	}
	lastUpdateId, err := employeeMapper.UpdateEmployee()
	if err != nil {
		return
	}
	employeeMapper, err = employeeMapper.GetEmployeeById(lastUpdateId)
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
func (e *Employee) DeleteEmployee(params *revel.Params) (employee Employee, err error) {
	emMapper := mappers.Employee{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}
	lastDeleteId, err := emMapper.DeleteEmployee(id)
	if err != nil {
		return
	}
	employee.Id = lastDeleteId
	return
}

func (e *Employee) AddEmployee(params *revel.Params) (employee Employee, err error) {
	employeeMapper := mappers.Employee{}
	err = json.Unmarshal(params.JSON, &employeeMapper)
	if err != nil {
		return
	}
	lastInsertedId, err := employeeMapper.AddEmployee()
	if err != nil {
		return
	}
	employeeMapper, err = employeeMapper.GetEmployeeById(lastInsertedId)
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

func (e *Employee) ResetPassEmployee(params *revel.Params) (employee Employee, err error) {
	emMapper := mappers.Employee{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	updatedRowId, err := emMapper.ResetPassEmployee(id)
	if err != nil {
		return
	}
	if id != updatedRowId {
		err = errors.New("not reset")
		return
	}
	return
}
