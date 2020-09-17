package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/entity"
	"myapp/app/providers"
	"strconv"
	"strings"
)

type Equipment struct {
	*revel.Controller
}
//
//func (c Equipment) Index() revel.Result {
//	return c.Render()
//}

//выдача товара сотруднику
func (c Equipment) DragToUser() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	renderInterface := app.RenderInterface{}
	equipmentModel := providers.EquipmentModel{}
	equipment := entity.Equipment{}

	err := json.Unmarshal(c.Params.JSON, &equipment)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := equipmentModel.DragToUser(app.DB, equipment)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)

}

//помещение товара на склад
func (c Equipment) DragToStore() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	DataEquipments := providers.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	var equipment entity.Equipment
	err := json.Unmarshal(c.Params.JSON, &equipment)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := DataEquipments.DragToStore(app.DB, equipment)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)

}

//получение оборудования,которое находится у определенного пользователя
func (c Equipment) GetEquipmentByUser() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	equipmentModel := providers.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	id := c.Params.Get("user")
	convId, err := strconv.Atoi(id)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := equipmentModel.GetEquipmentByUserId(app.DB, entity.Equipment{Id: convId})
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//получение оборудования,которое находится на складе
func (c Equipment) GetEquipmentsInStore() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	DataEquipments := providers.EquipmentModel{}
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

	if CheckPerm(c.Controller, "employee") || CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	DataEquipments := providers.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	token := c.Params.Get("token")
	session := Session[token]
	if session == "" {
		return c.Render()
	}
	splitSession := strings.Split(session, ":")
	var userId int
	var err error
	var tree []entity.FullTree
	if splitSession[1] == "employee" {
		userId, err = strconv.Atoi(splitSession[0])
		if err != nil {
			return c.Render()
		}
		tree, err = DataEquipments.GetEmployeeTree(app.DB, entity.Equipment{Fk_user1: userId})
	} else if splitSession[1] == "admin" {
		tree, err = DataEquipments.GetFullTree(app.DB)

	}
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = tree
	}
	return c.RenderJSON(renderInterface)
}

//получение всего оборудования
func (c Equipment) GetAllEquipments() revel.Result {

	if CheckPerm(c.Controller, "admin") || CheckPerm(c.Controller, "employee") {
	} else {
		return c.Render()
	}
	renderInterface := app.RenderInterface{}
	token := c.Params.Get("token")
	session := Session[token]
	if session == "" {
		return c.Render()
	}
	splitSession := strings.Split(session, ":")
	var userId int
	var equipments []entity.Equipment
	var err error
	var equipmentModel providers.EquipmentModel
	if splitSession[1] == "employee" {
		userId, err = strconv.Atoi(splitSession[0])
		if err != nil {
			return c.Render()
		}

		equipments, err = equipmentModel.GetEquipmentByUserId(app.DB, entity.Equipment{Id: userId})
	} else if splitSession[1] == "admin" {
		equipments, err = equipmentModel.GetAllEquipments(app.DB)
	}
		if err != nil {
			renderInterface.Error = err.Error()
		} else {
			renderInterface.Data = equipments
		}
		return c.RenderJSON(renderInterface)
}

//списать оборудование
func (c Equipment) WriteEquipment() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	DataEquipments := providers.EquipmentModel{}
	renderInterface := app.RenderInterface{}

	var equipment entity.Equipment
	err := json.Unmarshal(c.Params.JSON, &equipment)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := DataEquipments.WriteEquipment(app.DB, equipment)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//изменение оборудования
func (c Equipment) UpdateEquipment() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	equipmentModel := providers.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	var equipment entity.Equipment

	err := json.Unmarshal(c.Params.JSON, &equipment)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := equipmentModel.UpdateEquipment(app.DB, equipment)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//добавление оборудования
func (c Equipment) AddEquipment() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	equipmentModel := providers.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	var equipment entity.Equipment
	err := json.Unmarshal(c.Params.JSON, &equipment)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := equipmentModel.AddEquipment(app.DB, equipment)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}

//удаление оборудования
func (c Equipment) DeleteEquipment() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	equipmentModel := providers.EquipmentModel{}
	renderInterface := app.RenderInterface{}
	var equipment entity.Equipment
	err := json.Unmarshal(c.Params.JSON, &equipment)
	if err != nil {
		renderInterface.Error = err.Error()
		return c.RenderJSON(renderInterface)
	}
	result, err := equipmentModel.DeleteEquipment(app.DB, equipment)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result
	}
	return c.RenderJSON(renderInterface)
}
