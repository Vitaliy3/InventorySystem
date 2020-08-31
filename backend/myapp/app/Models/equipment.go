package models

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"myapp/app/mappers"
)

type RenderData struct {
	DataArray []EquipmentModel
	Data      EquipmentModel
	Error     error
	Tree      []FullTree
}

type EquipmentModel struct {
	Id              int        ` json:"id" `
	Fk_class        int        ` json:"class" `
	Fk_subclass     int        ` json:"subclass" `
	UserFIO         string     ` json:"user" `
	InventoryNumber string     ` json:"inventoryNumber" `
	EquipmentName   string     ` json:"name" `
	Status          string     ` json:"status" `
	ClassName       string     ` json:"value" `
	Data            []Subclass ` json:"data" `
}
type Subclass struct {
	Id           int    ` json:"subclass" `
	SubclassName string ` json:"value" `
}
type FullTree struct {
	Class string           `json:"class"`
	Value string           ` json:"value" `
	Tree  []EquipmentModel ` json:"data" `
}

func (e *EquipmentModel) UpdateEquipment(params *revel.Params) (render RenderData) {
	eqMapper := mappers.EquipmentTable{}
	err := json.Unmarshal(params.JSON, &eqMapper)
	fmt.Println("afterMarshall",eqMapper)
	result, err := eqMapper.UpdateEquipment()
	if err != nil {
		render.Error = err
		return
	}
	if result > 0 {
		row, _ := eqMapper.GetEquipmentById(eqMapper.Id)
		render.Data.Id = row.Id
		render.Data.Fk_class = row.Fk_parent
		render.Data.Fk_subclass = row.Fk_class
		render.Data.EquipmentName = row.EquipmentName
		render.Data.InventoryNumber = row.InventoryNumber
		render.Data.Status = getStatus(row.Status)
		return
	}
	return
}

//списать оборудование +
func (e *EquipmentModel) WriteEquipment(params *revel.Params) (render RenderData) {
	eqMapper := mappers.EquipmentTable{}
	var id int
	err := json.Unmarshal(params.JSON, &id)
	result, err := eqMapper.WriteEquipment(id)
	if err != nil {
		render.Error = err
		return
	}
	if result > 0 {
		row, _ := eqMapper.GetEquipmentById(id)
		if row.Status == 2 {
			status := "Списано"
			render.Data.Status = status
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
		temp.Fk_class = v.Fk_parent

		temp.Fk_subclass = v.Fk_class
		temp.InventoryNumber = v.InventoryNumber
		temp.EquipmentName = v.EquipmentName
		temp.Status = getStatus(v.Status)
		render.DataArray = append(render.DataArray, temp)
	}
	return render
}
func getStatus(status int) (newStatus string) {
	switch status {
	case 0:
		newStatus = "На складе"
	case 1:
		newStatus = "У сотрудника"
	case 2:
		newStatus = "Списано"
	}
	return
}
func (e *EquipmentModel) GetFullTree() (render RenderData) {
	eqMapper := mappers.EquipmentTable{}
	dbEqupments, err := eqMapper.GetFullTree()
	if err != nil {
		render.Error = err
		return
	}
	fullTree := render.DataArray
	for _, v := range dbEqupments {
		var temp EquipmentModel

		temp.ClassName = v.Class
		temp.Fk_class = v.Fk_parent
		subclass := Subclass{v.Fk_class, v.Subclass}
		temp.Data = append(temp.Data, subclass)
		find := false
		fmt.Println(fullTree)
		for i, q := range fullTree {
			if q.Fk_class == v.Fk_parent {
				find = true
				fullTree[i].Data = append(fullTree[i].Data, subclass)

				break
			}
		}
		if !find {
			fullTree = append(fullTree, temp)
		}
	}
	tree := FullTree{}
	tree.Value = "Все"
	tree.Class = "0"
	tree.Tree = fullTree
	render.Tree = append(render.Tree, tree)
	return
}

func (e *EquipmentModel) AddEquipment(c *revel.Params) (render RenderData) {
	eqMapper := mappers.EquipmentTable{}
	paramJson := c.JSON
	err := json.Unmarshal(paramJson, &eqMapper)
	if err != nil {
		fmt.Println("ErrAddEq:", err)
	}
	eqMapper.Status = 0
	fmt.Println("newEq", eqMapper)
	newEq, err := eqMapper.AddEquipment()
	if err != nil {
		fmt.Println("errAdd", err)
		render.Error = err
		return
	}
	if newEq > 0 {
		result, err := eqMapper.GetEquipmentById(newEq)
		render.Data.Id = result.Id
		render.Data.Fk_class = result.Fk_parent
		render.Data.Fk_subclass = result.Fk_class
		render.Data.EquipmentName = result.EquipmentName
		render.Data.InventoryNumber = result.InventoryNumber
		render.Data.Status = getStatus(result.Status)
		if err != nil {
			render.Error = err
			return
		}
	}
	fmt.Println("renderDataAdd", render.Data)
	return render
}

func (e *EquipmentModel) DeleteEquipment(params *revel.Params) (render RenderData) {
	eqMapper := mappers.EquipmentTable{}
	var id int
	err := json.Unmarshal(params.JSON, &id)
	if err != nil {
		fmt.Println("errUnmarshallDeleteEq", err)
		render.Error = err
		return
	}
	result, err := eqMapper.DeleteEquipment(id)
	if err != nil {
		render.Error = err
		return
	}
	if result > 0 {
		return render
	}
	return
}
