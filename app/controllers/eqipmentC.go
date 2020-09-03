package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) DragToUser() revel.Result {
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

func (c App) DragToStore() revel.Result {
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
func (c App) GetEquipmentOnUser() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	result, err := DataEquipments.GetEquipmentOnUser(app.DB, c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение оборудования,которое находится на складе
func (c App) GetEquipmentsInStore() revel.Result {
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
func (c App) GetFullTree() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	result, err := DataEquipments.GetFullTree(app.DB)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение всего оборудования
func (c App) GetAllEquipments() revel.Result {
	DataEquipments := models.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	result, err := DataEquipments.GetAllEquipments(app.DB)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//списать оборудование
func (c App) WriteEquipment() revel.Result {
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
func (c App) UpdateEquipment() revel.Result {
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
func (c App) AddEquipment() revel.Result {
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
func (c App) DeleteEquipment() revel.Result {
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
