package models

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"myapp/app/mappers"
)

type RenderDataEm struct {
	DataArray []Employee
	Data      Employee
	Error     error
}

type Employee struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Fk_role    int    `json:"fk_role"`
}

func (e *Employee) GetAllEmployees() (render RenderDataEm) {
	emMapper := mappers.Employee{}
	dbEmployees, err := emMapper.GetAllEmployees()
	if err != nil {
		render.Error = err
		return
	}
	var temp Employee
	for _, v := range dbEmployees {
		temp.Id = v.Id
		temp.Name = v.Name
		temp.Surname = v.Surname
		temp.Patronymic = v.Patronymic
		temp.Login = v.Login
		render.DataArray = append(render.DataArray, temp)
	}
	return
}
func (e *Employee) UpdateEmployee(params *revel.Params) (render RenderDataEm) {
	employeeMapper := mappers.Employee{}
	err := json.Unmarshal(params.JSON, &employeeMapper)
	fmt.Println("afterMarshall", employeeMapper)
	result, err := employeeMapper.UpdateEmployee()
	if err != nil {
		render.Error = err
		return
	}
	if result > 0 {
		employeeMapper, err = employeeMapper.GetEmployeeById(employeeMapper.Id)
		render.Data.Id = employeeMapper.Id
		render.Data.Name = employeeMapper.Name
		render.Data.Surname = employeeMapper.Surname
		render.Data.Patronymic = employeeMapper.Patronymic
		render.Data.Login = employeeMapper.Login
		if err != nil {
			render.Error = err
			return
		}
		return
	}
	return
}
func (e *Employee) DeleteEmployee(params *revel.Params) (render RenderDataEm) {
	emMapper := mappers.Employee{}
	var id int
	err := json.Unmarshal(params.JSON, &id)

	_, err = emMapper.DeleteEmployee(id)
	if err != nil {
		render.Error = err
		return
	}
	return
}

func (e *Employee) AddEmployee(params *revel.Params) (render RenderDataEm) {
	employeeMapper := mappers.Employee{}
	err := json.Unmarshal(params.JSON, &employeeMapper)
	fmt.Println("unmarshlEmoloyee", employeeMapper)

	result, err := employeeMapper.AddEmployee()
	if err != nil {
		fmt.Println("errAddEmployee:", err)
		render.Error = err
		return
	}
	if result > 0 {
		employeeMapper, err = employeeMapper.GetEmployeeById(int(result))
		render.Data.Id = employeeMapper.Id
		render.Data.Name = employeeMapper.Name
		render.Data.Surname = employeeMapper.Surname
		render.Data.Patronymic = employeeMapper.Patronymic
		render.Data.Login = employeeMapper.Login
		if err != nil {
			render.Error = err
			return
		}
		return
	}
	return
}

func (e *Employee) ResetPassEmployee(params *revel.Params) (render RenderDataEm) {
	emMapper := mappers.Employee{}
	var id int
	err := json.Unmarshal(params.JSON, &id)
	_, err = emMapper.ResetPassEmployee(id)
	if err != nil {
		render.Error = err
		return
	}
	return
}
