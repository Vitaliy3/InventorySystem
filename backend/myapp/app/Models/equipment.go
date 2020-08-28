package models

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"log"
	"myapp/app/mappers"
	"strconv"
)

type RenderData struct {
	DataArray []EquipmentModel
	Data      EquipmentModel
	Error     error
}

type EquipmentModel struct {
	Id              int    ` json:"id" `
	Fk_class        int    ` json:"class" `
	Fk_subclass     int    ` json:"sublcass" `
	UserFIO         string ` json:"user" `
	InventoryNumber string ` json:"inventoryNumber" `
	EquipmentName   string ` json:"name" `
	Status          string ` json:"status" `
}

//списать оборудование +
func (e *EquipmentModel) WriteEquipment(params *revel.Params) (render RenderData) {
	id := params.Query.Get("id")
	convId, _ := strconv.Atoi(id)
	eqMapper := mappers.EquipmentTable{}
	result, err := eqMapper.WriteEquipment(convId)
	if err != nil {
		render.Error = err
		return
	}
	if result > 0 {
		row, _ := eqMapper.GetEquipmentById(convId)
		if row.Status == 2 {
			status := "Списано"
			render.Data.Status = status
			fmt.Println("STATUS", status)
			return
		}
	}
	return
}

//получение всего оборудования +
func (e *EquipmentModel) GetAllEquipments() RenderData {
	eqMapper := mappers.EquipmentTable{}
	render := RenderData{}
	dbEqupments, err := eqMapper.GetAllEquipments()
	if err != nil {
		render.Error = err
		return render
	}
	var temp EquipmentModel
	for _, v := range dbEqupments {
		temp.Id = v.Id
		temp.Fk_class = v.Id
		temp.InventoryNumber = v.InventoryNumber
		temp.EquipmentName = v.EquipmentName
		switch v.Status {
		case 0:
			temp.Status = "На складе"
		case 1:
			temp.Status = "У сотрудника"
		case 2:
			temp.Status = "Списано"
		}
		render.DataArray = append(render.DataArray, temp)
	}
	return render
}

func (e *EquipmentModel) AddEquipment(c *revel.Params) (render RenderData) {
	eqMapper := mappers.EquipmentTable{}
	eqModel := EquipmentModel{}
	paramJson := c.JSON
	fmt.Println("JSON", paramJson)
	err := json.Unmarshal(paramJson, &eqModel)
	if err != nil {
		fmt.Println("Err:",err)
	}
	fmt.Println("AddEq", eqModel)
	newEq, err := eqMapper.AddEquipment()
	_, err = eqMapper.GetEquipmentById(int(newEq))
	if err != nil {
		render.Error = err
		return
	}
	return render
}

func (e *EquipmentModel) DeleteEquipment(params *revel.Params) (render RenderData) {
	eqModel := EquipmentModel{}
	rawJson := params.JSON
	err := json.Unmarshal(rawJson, &eqModel)
	if err != nil {
		fmt.Println("ERR:",err)
	}
	eqMapper := mappers.EquipmentTable{}
	fmt.Println("ID:",eqModel.Id)
	result, err := eqMapper.DeleteEquipment(eqModel.Id)
	if err != nil {
		log.Println(err)
	}
	if result > 0 {
		return render
	}
	return
}
