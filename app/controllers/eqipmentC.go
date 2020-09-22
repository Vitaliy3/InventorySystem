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
	renderInterface app.RenderInterface
	equipmentModel  providers.EquipmentModel
	equipment       entity.Equipment
}

//выдача товара сотруднику
func (c Equipment) DragToUser() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.equipmentModel.DragToUser(app.Db, c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//перемещение товара на склад
func (c Equipment) DragToStore() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.equipmentModel.DragToStore(app.Db, c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//получение оборудования,которое находится у определенного пользователя
func (c Equipment) GetByUserId() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	id := c.Params.Get("user")
	convId, err := strconv.Atoi(id)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.equipmentModel.GetByUserId(app.Db, entity.Equipment{Id: convId})
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//получение оборудования,которое находится на складе
func (c Equipment) GetInStore() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	result, err := c.equipmentModel.GetInStore(app.Db)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//получение полного дерева учета оборудования
func (c Equipment) GetTree() revel.Result {
	if CheckPerm(c.Controller, "employee") || CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	var (
		userId int
		err    error
		tree   []entity.FullTree
	)
	token := c.Params.Get("token") //получение токена
	session := Session[token]
	if session == "" {
		return c.Render()
	}

	splitSession := strings.Split(session, ":")
	if splitSession[1] == "employee" {
		userId, err = strconv.Atoi(splitSession[0])
		if err != nil {
			return c.Render()
		}

		tree, err = c.equipmentModel.GetTreeByUserId(app.Db, entity.Equipment{Fk_userI: userId}) //получение дерева для конкретного пользователя
	} else if splitSession[1] == "admin" {
		tree, err = c.equipmentModel.GetTree(app.Db)
	}
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = tree
	}
	return c.RenderJSON(c.renderInterface)
}

//получение всего оборудования
func (c Equipment) GetAll() revel.Result {
	if CheckPerm(c.Controller, "admin") || CheckPerm(c.Controller, "employee") {
	} else {
		return c.Render()
	}

	var (
		userId     int
		equipments []entity.Equipment
		err        error
	)

	token := c.Params.Get("token")
	session := Session[token]
	if session == "" {
		return c.Render()
	}

	splitSession := strings.Split(session, ":")
	if splitSession[1] == "employee" {
		userId, err = strconv.Atoi(splitSession[0])
		if err != nil {
			return c.Render()
		}
		equipments, err = c.equipmentModel.GetByUserId(app.Db, entity.Equipment{Id: userId})
	} else if splitSession[1] == "admin" {
		equipments, err = c.equipmentModel.GetAll(app.Db)
	}
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = equipments
	}
	return c.RenderJSON(c.renderInterface)
}

//списать оборудование
func (c Equipment) Write() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.equipmentModel.Write(app.Db, c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//изменение оборудования
func (c Equipment) Update() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.equipmentModel.Update(app.Db, c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//добавление оборудования
func (c Equipment) Add() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.equipmentModel.Add(app.Db, c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}

//удаление оборудования
func (c Equipment) Delete() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	err := json.Unmarshal(c.Params.JSON, &c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
		return c.RenderJSON(c.renderInterface)
	}

	result, err := c.equipmentModel.Delete(app.Db, c.equipment)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}
