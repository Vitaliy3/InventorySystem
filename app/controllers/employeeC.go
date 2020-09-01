package controllers

import (
	"github.com/revel/revel"
	"myapp/app"
)

type Employee struct {
	*revel.Controller
}

//добавление сотрудников
func (c Employee) AddEmployee() revel.Result {
	DataEmployee := app.RenderDataEmployee{}
	result, err := DataEmployee.Data.AddEmployee(c.Params)
	if err != nil {
		DataEmployee.Error = err.Error()
	} else {
		DataEmployee.Data = result
	}
	return c.RenderJSON(DataEmployee)
}

//получение всех сотрудников
func (c Employee) GetAllEmployees() revel.Result {
	DataEmployee := app.RenderDataEmployee{}
	result, err := DataEmployee.Data.GetAllEmployees()
	if err != nil {
		DataEmployee.Error = err.Error()
	} else {
		DataEmployee.DataArray = result
	}
	return c.RenderJSON(DataEmployee)
}

//изменение данных о сотруднике
func (c Employee) UpdateEmployee() revel.Result {
	DataEmployee := app.RenderDataEmployee{}
	result, err := DataEmployee.Data.UpdateEmployee(c.Params)
	if err != nil {
		DataEmployee.Error = err.Error()
	} else {
		DataEmployee.Data = result
	}
	return c.RenderJSON(DataEmployee)
}

//сброс пароля у сотрудника
func (c Employee) ResetPassEmployee() revel.Result {
	DataEmployee := app.RenderDataEmployee{}
	result, err := DataEmployee.Data.ResetPassEmployee(c.Params)
	if err != nil {
		DataEmployee.Error = err.Error()
	} else {
		DataEmployee.Data = result
	}
	return c.RenderJSON(DataEmployee)
}

//удаление сотрудника
func (c Employee) DeleteEmployee() revel.Result {
	DataEmployee := app.RenderDataEmployee{}
	result, err := DataEmployee.Data.DeleteEmployee(c.Params)
	if err != nil {
		DataEmployee.Error = err.Error()
	} else {
		DataEmployee.Data = result
	}
	return c.RenderJSON(DataEmployee)
}
