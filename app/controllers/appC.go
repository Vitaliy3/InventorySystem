package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

//получение полного дерева учета оборудования
func (c App) GetFullTree() revel.Result {
	DataEquipments:=app.RenderDataEquipments{}
	result, err := DataEquipments.Data.GetFullTree()
	if err != nil {
		DataEquipments.Error = err.Error()
	} else {
		DataEquipments.Tree = result
	}
	return c.RenderJSON(DataEquipments)
}

//получение всего оборудования
func (c App) GetAllEquipments() revel.Result {
	DataEquipments:=app.RenderDataEquipments{}

	result, err := DataEquipments.Data.GetAllEquipments()
	if err != nil {
		DataEquipments.Error = err.Error()
	} else {
		DataEquipments.DataArray = result
	}
	return c.RenderJSON(DataEquipments)
}

//списать оборудование
func (c App) WriteEquipment() revel.Result {
	DataEquipments:=app.RenderDataEquipments{}

	result, err := DataEquipments.Data.WriteEquipment(c.Params)
	if err != nil {
		DataEquipments.Error = err.Error()
	} else {
		DataEquipments.Data = result
	}
	return c.RenderJSON(DataEquipments)
}

//изменение оборудования
func (c App) UpdateEquipment() revel.Result {
	DataEquipments:=app.RenderDataEquipments{}

	result, err := DataEquipments.Data.UpdateEquipment(c.Params)
	if err != nil {
		DataEquipments.Error = err.Error()
	} else {
		DataEquipments.Data = result
	}
	return c.RenderJSON(DataEquipments)
}

//добавление оборудования
func (c App) AddEquipment() revel.Result {
	DataEquipments:=app.RenderDataEquipments{}

	result, err := DataEquipments.Data.AddEquipment(c.Params)
	fmt.Println("RESULT", result)
	if err != nil {
		DataEquipments.Error = err.Error()
	} else {
		DataEquipments.Data = result
	}
	return c.RenderJSON(DataEquipments)
}

//удаление оборудования
func (c App) DeleteEquipment() revel.Result {
	DataEquipments:=app.RenderDataEquipments{}

	result, err := DataEquipments.Data.DeleteEquipment(c.Params)
	if err != nil {
		DataEquipments.Error = err.Error()
	} else {
		DataEquipments.Data = result
	}
	return c.RenderJSON(DataEquipments)
}
