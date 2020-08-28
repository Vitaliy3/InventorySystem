package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetAllEquipments() revel.Result {
	equipment := models.EquipmentModel{}
	result := equipment.GetAllEquipments()
	return c.RenderJSON(result)
}

//списать оборудование +
func (c App) WriteEquipment() revel.Result {
	equipment := models.EquipmentModel{}
	result := equipment.WriteEquipment(c.Params)
	return c.RenderJSON(result)
}

func (c App) AddEquipment() revel.Result {
	equipment := models.EquipmentModel{}
	result := equipment.AddEquipment(c.Params)
	return c.RenderJSON(result)
}
func (c App) DeleteEquipment() revel.Result {
	equipment := models.EquipmentModel{}
	result := equipment.DeleteEquipment(c.Params)
	return c.RenderJSON(result)
}
