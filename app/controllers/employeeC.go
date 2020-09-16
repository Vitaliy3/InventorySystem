package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/entity"
	"myapp/app/models"
)

type User struct {
	*revel.Controller
}

//добавление сотрудников
func (c User) AddEmployee() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	renderInterface := app.RenderInterface{}
	var employee entity.Employee
	employeeModel := models.Employee{}
	err := json.Unmarshal(c.Params.JSON, &employee)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := employeeModel.AddEmployee(app.DB, employee)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение всех сотрудников
func (c User) GetAllEmployees() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	employeeModel := models.Employee{}
	renderInterface := app.RenderInterface{}
	result, err := employeeModel.GetAllEmployees(app.DB)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//изменение данных о сотруднике
func (c User) UpdateEmployee() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	employeeModel := models.Employee{}
	renderInterface := app.RenderInterface{}
	employee := entity.Employee{}

	err := json.Unmarshal(c.Params.JSON, &employee)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := employeeModel.UpdateEmployee(app.DB, employee)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//сброс пароля у сотрудника
func (c User) ResetPassEmployee() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	DataEmployee := models.Employee{}
	renderInterface := app.RenderInterface{}
	var employee entity.Employee

	err := json.Unmarshal(c.Params.JSON, &employee)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := DataEmployee.ResetPassEmployee(app.DB, employee)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//удаление сотрудника
func (c User) DeleteEmployee() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	DataEmployee := models.Employee{}
	renderInterface := app.RenderInterface{}
	var employee entity.Employee

	err := json.Unmarshal(c.Params.JSON, &employee)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := DataEmployee.DeleteEmployee(app.DB, employee)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}
