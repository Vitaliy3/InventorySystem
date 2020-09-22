package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/entity"
	"myapp/app/providers"
)

type User struct {
	*revel.Controller
	renderInterface app.RenderInterface
	employee        entity.Employee
	employeeModel   providers.Employee
}

//добавление сотрудников
func (c User) Add() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.employeeModel.Add(app.Db, c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//получение всех сотрудников
func (c User) GetAll() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	result, err := c.employeeModel.GetAll(app.Db)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//изменение данных о сотруднике
func (c User) Update() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.employeeModel.Update(app.Db, c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//сброс пароля у сотрудника
func (c User) ResetPassword() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.employeeModel.ResetPassword(app.Db, c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//удаление сотрудника
func (c User) Delete() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.employeeModel.Delete(app.Db, c.employee)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}
