package models

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"myapp/app/mappers"
)

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
	Class      string           `json:"class"`
	Value      string           ` json:"value" `
	EquipModel []EquipmentModel ` json:"data" `
}

//изменение оборудования
func (e *EquipmentModel) UpdateEquipment(params *revel.Params) (equip EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	eqModel := EquipmentModel{}
	err = json.Unmarshal(params.JSON, &eqModel)
	if err != nil {
		return
	}
	eqMapper.Id=eqModel.Id
	eqMapper.EquipmentName=eqModel.EquipmentName
	eqMapper.InventoryNumber=eqModel.InventoryNumber
	lastInsertedId, err := eqMapper.UpdateEquipment()
	if err != nil {
		return
	}
	row, err := eqMapper.GetEquipmentById(lastInsertedId)
	if err != nil {
		return
	}
	fmt.Println("do", row)
	equip.Id = row.Id
	equip.Fk_class = row.Fk_parent
	equip.Fk_subclass = row.Fk_class
	equip.EquipmentName = row.EquipmentName
	equip.InventoryNumber = row.InventoryNumber
	equip.Status = getStatus(row.Status)
	fmt.Println("after", equip)

	return
}

//списать оборудование
func (e *EquipmentModel) WriteEquipment(params *revel.Params) (equip EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}
	result, err := eqMapper.WriteEquipment(id)
	if err != nil {
		return
	}
	row, err := eqMapper.GetEquipmentById(result)
	if err != nil {
		return
	}

	if row.Status == 2 {
		equip.Id = row.Id
		equip.Status = "Списано"
		return
	}
	return
}

type RenderData struct {
	DataArray []EquipmentModel
	Data      EquipmentModel
	Error     error
	Tree      []FullTree
}

//получение всего оборудования
func (e *EquipmentModel) GetAllEquipments() (equipArray []EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	dbEquip, err := eqMapper.GetAllEquipments()
	if err != nil {
		return
	}
	var temp EquipmentModel
	for _, v := range dbEquip {
		temp.Id = v.Id
		temp.Fk_class = v.Fk_parent
		temp.Fk_subclass = v.Fk_class
		temp.InventoryNumber = v.InventoryNumber
		temp.EquipmentName = v.EquipmentName
		temp.Status = getStatus(v.Status)
		equipArray = append(equipArray, temp)
	}
	return
}

//определение статуса оборудования
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

//получение полного дерева учета оборудования
func (e *EquipmentModel) GetFullTree() (fullTree []FullTree, err error) {
	eqMapper := mappers.EquipmentTable{}
	dbEqupments, err := eqMapper.GetFullTree()
	if err != nil {
		return
	}
	var trees []EquipmentModel
	for _, v := range dbEqupments {
		var temp EquipmentModel
		temp.ClassName = v.Class
		temp.Fk_class = v.Fk_parent
		subclass := Subclass{v.Fk_class, v.Subclass}
		temp.Data = append(temp.Data, subclass)
		find := false
		fmt.Println(trees)
		for i, q := range trees { //поиск уже добавленной записи в дерево
			if q.Fk_class == v.Fk_parent {
				find = true
				trees[i].Data = append(trees[i].Data, subclass)
				break
			}
		}
		if !find {
			trees = append(trees, temp)
		}
	}
	tree := FullTree{}
	tree.Value = "Все"
	tree.Class = "0"
	tree.EquipModel = trees
	fullTree = append(fullTree, tree)
	return
}

//добавление оборудования
func (e *EquipmentModel) AddEquipment(c *revel.Params) (eqData EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	paramJson := c.JSON
	err = json.Unmarshal(paramJson, &eqMapper)
	if err != nil {
		return
	}
	eqMapper.Status = 0
	fmt.Println("newEq", eqMapper)
	lastInsertedId, err := eqMapper.AddEquipment()
	if err != nil {
		return
	}
	result, err := eqMapper.GetEquipmentById(lastInsertedId)
	if err != nil {
		return
	}
	eqData.Id = result.Id
	eqData.Fk_class = result.Fk_parent
	eqData.Fk_subclass = result.Fk_class
	eqData.EquipmentName = result.EquipmentName
	eqData.InventoryNumber = result.InventoryNumber
	eqData.Status = getStatus(result.Status)
	return
}

//удаление оборудования
func (e *EquipmentModel) DeleteEquipment(params *revel.Params) (data EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}
	rowsAffected, err := eqMapper.DeleteEquipment(id)
	if err != nil {
		return
	}
	if rowsAffected > 0 {
		data.Id = id
	}
	return
}
