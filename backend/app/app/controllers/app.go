package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type Equipment struct {
	Id              string ` json:"id" `
	Class           string ` json:"class" `
	Subclass        string` json:"subclass" `
	Name            string` json:"name" `
	User            string` json:"user" `
	Status          string` json:"status" `
	InventoryNumber string` json:"inventoryNumber" `
}

func (c App) Last() revel.Result {
	equipment := Equipment{"1", "1", "1", "item1", "user", "1", "001"}
	return c.RenderJSON(equipment)
}


func (c App) GetEquipment() revel.Result {
	fmt.Println(c.Params.Query.Get("user"))
	equipment := Equipment{"1", "1", "1", "item1", "user", "1", "001"}
	return c.RenderJSON(equipment)
}