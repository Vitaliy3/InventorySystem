package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/models"
)

type Equipment struct {
	*revel.Controller
}

func (c Equipment) Index() revel.Result {
	return c.Render()
}

func (c Equipment) DragToUser() revel.Result {
	renderInterface := app.RenderInterface{}
	equipModel := models.EquipmentModel{}
	result, err := equipModel.DragToUser(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)

}

func (c Equipment) DragToStore() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	result, err := DataEquipments.DragToStore(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)

}

//получение оборудования,которое находится у определенного пользователя
func (c Equipment) GetEquipmentByUser() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}

	result, err := DataEquipments.GetEquipmentByUser(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение оборудования,которое находится на складе
func (c Equipment) GetEquipmentsInStore() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	result, err := DataEquipments.GetEquipmentsInStore(app.DB)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение полного дерева учета оборудования
func (c Equipment) GetFullTree() revel.Result {

	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	result, err := DataEquipments.GetFullTree(app.DB, c.Params, Session)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение всего оборудования
func (c Equipment) GetAllEquipments() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}

	result, err := DataEquipments.GetAllEquipments(app.DB, c.Params, Session)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	fmt.Println("RENDERDATA", renderInterface.Data)
	return c.RenderJSON(renderInterface)
}

//списать оборудование
func (c Equipment) WriteEquipment() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}

	result, err := DataEquipments.WriteEquipment(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//изменение оборудования
func (c Equipment) UpdateEquipment() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}

	result, err := DataEquipments.UpdateEquipment(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//добавление оборудования
func (c Equipment) AddEquipment() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}

	result, err := DataEquipments.AddEquipment(app.DB, c.Params)
	fmt.Println("RESULT", result)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//удаление оборудования
func (c Equipment) DeleteEquipment() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}

	result, err := DataEquipments.DeleteEquipment(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}
