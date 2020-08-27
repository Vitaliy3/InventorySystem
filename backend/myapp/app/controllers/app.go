package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app/Models"
	"strconv"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetAllEquipments() revel.Result {
	newModel := Models.EquipmentModel{}
	result := newModel.GetAllEquipments()
	fmt.Println(result)
	return c.RenderJSON(result)
}

func (c App) WriteEquipment() revel.Result {
	newModel := Models.EquipmentModel{}
	id := c.Params.Query.Get("id")
	convId, _ := strconv.Atoi(id)
	status := newModel.WriteEquipmentModel(convId)
	return c.RenderJSON(status)
}
type data struct{
	id string `json:"id"`
	name string `json:"name"`
}
func (c App) AddEquipment() revel.Result {
	data := data{}
	fmt.Println(*c.Params)
	_ = c.Params.BindJSON(&data)
	fmt.Println("POST:",data)
	return c.RenderJSON(data)
}
