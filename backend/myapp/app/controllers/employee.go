package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
)

func (c App) AddEmployee() revel.Result {
	employee := models.Employee{}
	result := employee.AddEmployee(c.Params)
	return c.RenderJSON(result)
}
func (c App) GetAllEmployees() revel.Result {
	employee := models.Employee{}
	result := employee.GetAllEmployees()
	return c.RenderJSON(result)
}

func (c App) UpdateEmployee() revel.Result {
	employee := models.Employee{}
	result := employee.UpdateEmployee(c.Params)
	return c.RenderJSON(result)
}
func (c App) ResetPassEmployee() revel.Result {
	employee := models.Employee{}
	result := employee.ResetPassEmployee(c.Params)
	return c.RenderJSON(result)
}
func (c App) DeleteEmployee() revel.Result {
	employee := models.Employee{}
	result := employee.DeleteEmployee(c.Params)
	return c.RenderJSON(result)
}
