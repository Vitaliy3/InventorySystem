package controllers

import (
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/models"
)

type User struct {
	*revel.Controller
}

//добавление сотрудников
func (c User) AddEmployee() revel.Result {
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
func (c User) GetAllEmployees() revel.Result {
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
func (c User) UpdateEmployee() revel.Result {
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
func (c User) ResetPassEmployee() revel.Result {
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
func (c User) DeleteEmployee() revel.Result {
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
