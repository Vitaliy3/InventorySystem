package controllers

import (
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/models"
)

type Employee struct {
	*revel.Controller
}

//добавление сотрудников
func (c Employee) AddEmployee() revel.Result {
	DataEmployee := models.Employee{}
	renderInterface := app.RenderInterface{}
	result, err := DataEmployee.AddEmployee(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение всех сотрудников
func (c Employee) GetAllEmployees() revel.Result {
	DataEmployee := models.Employee{}
	renderInterface := app.RenderInterface{}
	result, err := DataEmployee.GetAllEmployees(app.DB)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//изменение данных о сотруднике
func (c Employee) UpdateEmployee() revel.Result {
	DataEmployee := models.Employee{}
	renderInterface := app.RenderInterface{}
	result, err := DataEmployee.UpdateEmployee(app.DB,c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//сброс пароля у сотрудника
func (c Employee) ResetPassEmployee() revel.Result {
	DataEmployee := models.Employee{}
	renderInterface := app.RenderInterface{}
	result, err := DataEmployee.ResetPassEmployee(app.DB,c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//удаление сотрудника
func (c Employee) DeleteEmployee() revel.Result {
	DataEmployee := models.Employee{}
	renderInterface := app.RenderInterface{}
	result, err := DataEmployee.DeleteEmployee(app.DB,c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}
