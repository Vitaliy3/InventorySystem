package models

import (
	"database/sql"
	"encoding/json"
	"github.com/revel/revel"
	"myapp/app/mappers"
	"strconv"
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
	Fk_user         int        `json:"fk_user"`
	Data            []Subclass ` json:"data" `
	Open            bool       `json:"open"`
}
type Subclass struct {
	Id           int    ` json:"subclass" `
	SubclassName string ` json:"value" `
}
type FullTree struct {
	Class      string           `json:"class"`
	Value      string           ` json:"value" `
	EquipModel []EquipmentModel ` json:"data" `
	Open       bool             `json:"open"`
}

//перемещение оборудования со склада пользователю
func (e *EquipmentModel) DragToUser(DB *sql.DB, params *revel.Params) (equip EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	eqMod := EquipmentModel{}
	err = json.Unmarshal(params.JSON, &eqMod)
	if err != nil {
		return
	}
	updatedRowId, err := eqMapper.DragToUser(DB, eqMod.Id, eqMod.Fk_user)
	if err != nil {
		return
	}
	row, err := eqMapper.GetEquipmentById(DB, updatedRowId)
	if err != nil {
		return
	}
	equip.Id = row.Id
	equip.Fk_class = row.Fk_parent
	equip.Fk_subclass = row.Fk_class
	equip.EquipmentName = row.EquipmentName
	equip.InventoryNumber = row.InventoryNumber
	equip.Status = getStatus(row.Status)
	return
}

//перемещение оборудования от сотрудника на склад
func (e *EquipmentModel) DragToStore(DB *sql.DB, params *revel.Params) (equip EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}

	updatedRowId, err := eqMapper.DragToStore(DB, id)
	if err != nil {
		return
	}
	row, err := eqMapper.GetEquipmentById(DB, updatedRowId)
	if err != nil {
		return
	}
	equip.Id = row.Id
	equip.Fk_class = row.Fk_parent
	equip.Fk_subclass = row.Fk_class
	equip.EquipmentName = row.EquipmentName
	equip.InventoryNumber = row.InventoryNumber
	equip.Status = getStatus(row.Status)
	return
}

//изменение оборудования
func (e *EquipmentModel) UpdateEquipment(DB *sql.DB, params *revel.Params) (equip EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	eqModel := EquipmentModel{}
	err = json.Unmarshal(params.JSON, &eqModel)
	if err != nil {
		return
	}
	eqMapper.Id = eqModel.Id
	eqMapper.EquipmentName = eqModel.EquipmentName
	eqMapper.InventoryNumber = eqModel.InventoryNumber
	lastInsertedId, err := eqMapper.UpdateEquipment(DB)
	if err != nil {
		return
	}
	row, err := eqMapper.GetEquipmentById(DB, lastInsertedId)
	if err != nil {
		return
	}
	equip.Id = row.Id
	equip.Fk_class = row.Fk_parent
	equip.Fk_subclass = row.Fk_class
	equip.EquipmentName = row.EquipmentName
	equip.InventoryNumber = row.InventoryNumber
	equip.Status = getStatus(row.Status)

	return
}

//списать оборудование
func (e *EquipmentModel) WriteEquipment(DB *sql.DB, params *revel.Params) (equip EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}
	result, err := eqMapper.WriteEquipment(DB, id)
	if err != nil {
		return
	}
	row, err := eqMapper.GetEquipmentById(DB, result)
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

//получение продуктов, которые находятся на складе
func (e *EquipmentModel) GetEquipmentsInStore(DB *sql.DB) (equipArray []EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	result, err := eqMapper.GetEquipmentsInStore(DB)
	if err != nil {
		return
	}
	var temp EquipmentModel
	for _, v := range result {
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

//получение всего оборудования
func (e *EquipmentModel) GetAllEquipments(DB *sql.DB) (equipArray []EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	employeeMapper := mappers.Employee{}
	dbEquip, err := eqMapper.GetAllEquipments(DB)
	if err != nil {
		return
	}
	employees, err := employeeMapper.GetAllEmployees(DB)

	var temp EquipmentModel
	for _, v := range dbEquip {
		temp.Id = v.Id
		temp.Fk_class = v.Fk_parent
		temp.Fk_subclass = v.Fk_class
		temp.InventoryNumber = v.InventoryNumber
		temp.EquipmentName = v.EquipmentName
		temp.Status = getStatus(v.Status)
		for _, m := range employees {
			if int(v.Fk_user.Int64) == m.Id {
				temp.UserFIO = m.Name + " " + m.Surname + " " + m.Patronymic
			}
			}
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
func (e *EquipmentModel) GetFullTree(DB *sql.DB) (fullTree []FullTree, err error) {
	eqMapper := mappers.EquipmentTable{}
	dbEqupments, err := eqMapper.GetFullTree(DB)
	if err != nil {
		return
	}
	var trees []EquipmentModel
	for _, v := range dbEqupments {
		var temp EquipmentModel
		temp.ClassName = v.Class
		temp.Fk_class = v.Fk_parent
		temp.Open = true
		subclass := Subclass{v.Fk_class, v.Subclass}
		temp.Data = append(temp.Data, subclass)
		find := false
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
	tree.Open = true
	tree.EquipModel = trees
	fullTree = append(fullTree, tree)
	return
}

//добавление оборудования
func (e *EquipmentModel) AddEquipment(DB *sql.DB, c *revel.Params) (eqData EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	paramJson := c.JSON
	err = json.Unmarshal(paramJson, &eqMapper)
	if err != nil {
		return
	}
	eqMapper.Status = 0
	lastInsertedId, err := eqMapper.AddEquipment(DB)
	if err != nil {
		return
	}
	result, err := eqMapper.GetEquipmentById(DB, lastInsertedId)
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

func (e *EquipmentModel) GetEquipmentOnUser(DB *sql.DB, params *revel.Params) (equipArray []EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	userId := params.Get("user")
	convUserId, err := strconv.Atoi(userId)
	dbEquip, err := eqMapper.GetEquipmentsByUser(DB, convUserId)
	if err != nil {
		return
	}
	var temp EquipmentModel
	for _, v := range dbEquip {
		temp.Id = v.Id
		temp.Fk_user = convUserId
		temp.Fk_class = v.Fk_parent
		temp.Fk_subclass = v.Fk_class
		temp.Status = "На складе" //для переноса на склад от сотрудника
		temp.InventoryNumber = v.InventoryNumber
		temp.EquipmentName = v.EquipmentName
		equipArray = append(equipArray, temp)
	}
	return
}

//удаление оборудования
func (e *EquipmentModel) DeleteEquipment(DB *sql.DB, params *revel.Params) (data EquipmentModel, err error) {
	eqMapper := mappers.EquipmentTable{}
	var id int
	err = json.Unmarshal(params.JSON, &id)
	if err != nil {
		return
	}
	rowsAffected, err := eqMapper.DeleteEquipment(DB, id)
	if err != nil {
		return
	}
	if rowsAffected > 0 {
		data.Id = id
	}
	return
}
