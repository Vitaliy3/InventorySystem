package models

import (
	"database/sql"
	"myapp/app/entity"
	"myapp/app/mappers"
)

type EquipmentModel struct {
	entity.Equipment
}
type FullTree struct {
	entity.FullTree
}

//перемещение оборудования со склада пользователю
func (e *EquipmentModel) DragToUser(DB *sql.DB, equipment entity.Equipment) (equip entity.Equipment, err error) {
	equipmentMapper := mappers.Equipment{}
	updatedRowId, err := equipmentMapper.DragEquipmentToUser(DB, equipment)
	if err != nil {
		return
	}
	equip, err = equipmentMapper.GetEquipmentById(DB, updatedRowId)
	if err != nil {
		return
	}
	_, err = equipmentMapper.NewEvent(DB, entity.InventoryEvent{Fk_userI: equipment.Fk_user1, Fk_equipment: equipment.Id, ActionEvent: "Выдача сотруднику"})
	if err != nil {
		return
	}
	equip.Status = getStatus(equipment.StatusI)
	return
}

//перемещение оборудования от сотрудника на склад
func (e *EquipmentModel) DragToStore(DB *sql.DB, equipment entity.Equipment) (equip entity.Equipment, err error) {
	equipmentMapper := mappers.Equipment{}

	updatedRowId, err := equipmentMapper.DragEquipmentToStore(DB, equipment)
	if err != nil {
		return
	}
	equip, err = equipmentMapper.GetEquipmentById(DB, updatedRowId)
	if err != nil {
		return
	}
	_, err = equipmentMapper.NewEvent(DB, entity.InventoryEvent{Fk_userI:0, Fk_equipment: equip.Id, ActionEvent: "Перемещение на склад"})
	if err != nil {
		return
	}
	equip.Status = getStatus(equip.StatusI)
	return
}

//изменение оборудования
func (e *EquipmentModel) UpdateEquipment(DB *sql.DB, equipment entity.Equipment) (equip entity.Equipment, err error) {
	eqMapper := mappers.Equipment{}
	employee := mappers.Employee{}

	lastInsertedId, err := eqMapper.UpdateEquipment(DB, equipment)
	if err != nil {
		return
	}
	row, err := eqMapper.GetEquipmentById(DB, lastInsertedId)
	if err != nil {
		return
	}
	user, err := employee.GetEmployeeById(DB, int(row.Fk_user.Int64))
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	equip.Id = row.Id
	equip.Fk_class = row.Fk_parent
	if user.Name != "" {
		equip.UserFIO = user.Name + " " + user.Surname + " " + user.Patronymic
	} else {
		equip.UserFIO = "Отсутсвует"
	}
	equip.Fk_class = row.Fk_class
	equip.EquipmentName = row.EquipmentName
	equip.InventoryNumber = row.InventoryNumber
	equip.Status = getStatus(row.StatusI)
	return
}

//списать оборудование
func (e *EquipmentModel) WriteEquipment(DB *sql.DB, equipment entity.Equipment) (equipmentOut entity.Equipment, err error) {
	eqMapper := mappers.Equipment{}

	updatedId, err := eqMapper.WriteEquipment(DB, equipment)
	if err != nil {

		return
	}
	equipmentOut, err = eqMapper.GetEquipmentById(DB, updatedId)
	if err != nil {

		return
	}
	if equipmentOut.StatusI == 2 {
		equipmentOut.Status = "Списано"
	}
	return
}

//получение продуктов, которые находятся на складе
func (e *EquipmentModel) GetEquipmentsInStore(DB *sql.DB) (equipments []entity.Equipment, err error) {
	equipmentMapper := mappers.Equipment{}
	equipments, err = equipmentMapper.GetEquipmentsInStore(DB)
	for i, _ := range equipments {
		equipments[i].Status = getStatus(equipments[i].StatusI)
	}
	return
}

//получение всего оборудования
func (e *EquipmentModel) GetAllEquipments(DB *sql.DB) (equipments []entity.Equipment, err error) {
	equipmentMapper := mappers.Equipment{}
	employeeMapper := mappers.Employee{}
	equipments, err = equipmentMapper.GetAllEquipments(DB)
	employees, err := employeeMapper.GetAllEmployees(DB)
	for i, _ := range equipments {
		equipments[i].Status = getStatus(equipments[i].StatusI)
		for _, m := range employees {
			if int(equipments[i].Fk_user.Int64) == m.Id {
				equipments[i].UserFIO = m.Name + " " + m.Surname + " " + m.Patronymic
			}
		}
		if equipments[i].UserFIO == "" {
			equipments[i].UserFIO = "Отсутствует"
		}
	}
	return
}

func (e *EquipmentModel) GetEmployeeEquipments(DB *sql.DB, equipment entity.Equipment) (equipments []entity.Equipment, err error) {
	equipmentMapper := mappers.Equipment{}
	employeeMapper := mappers.Employee{}
	equipments, err = equipmentMapper.GetEquipmentsByUserId(DB, equipment)
	employees, err := employeeMapper.GetAllEmployees(DB)
	for i, _ := range equipments {
		equipments[i].Status = getStatus(equipments[i].StatusI)
		for _, m := range employees {
			if int(equipments[i].Fk_user.Int64) == m.Id {
				equipments[i].UserFIO = m.Name + " " + m.Surname + " " + m.Patronymic
			}
		}
		if equipments[i].UserFIO == "" {
			equipments[i].UserFIO = "Отсутствует"
		}
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
func (e *EquipmentModel) GetEmployeeTree(DB *sql.DB, equipment entity.Equipment) (fullTree []entity.FullTree, err error) {
	equipmentMapper := mappers.Equipment{}
	var equipments []entity.Equipment
	equipments, err = equipmentMapper.GetEmployeeTreeById(DB,equipment)
	if err != nil {
		return
	}
	var trees []entity.Equipment
	for _, v := range equipments {
		var temp entity.Equipment
		temp.ClassName = v.Class
		temp.Fk_parent = v.Fk_parent
		temp.Open = true
		subclass := entity.Subclass{v.Fk_class, v.Subclass}
		temp.Data = append(temp.Data, subclass)
		find := false
		for i, q := range trees { //поиск уже добавленной записи в дерево
			if q.Fk_parent == v.Fk_parent {
				find = true
				trees[i].Data = append(trees[i].Data, subclass)
				break
			}
		}
		if !find {
			trees = append(trees, temp)
		}
	}
	tree := entity.FullTree{}
	tree.Value = "Все"
	tree.Id = 0
	tree.Open = true
	tree.Equipment = trees
	fullTree = append(fullTree, tree)
	return
}

func (e *EquipmentModel) GetFullTree(DB *sql.DB) (fullTree []entity.FullTree, err error) {
	equipmentMapper := mappers.Equipment{}
	var equipments []entity.Equipment
	equipments, err = equipmentMapper.GetFullTree(DB)
	if err != nil {
		return
	}
	var trees []entity.Equipment
	for _, v := range equipments {
		var temp entity.Equipment
		temp.ClassName = v.Class
		temp.Fk_parent = v.Fk_parent
		temp.Open = true
		subclass := entity.Subclass{v.Fk_class, v.Subclass}
		temp.Data = append(temp.Data, subclass)
		find := false
		for i, q := range trees { //поиск уже добавленной записи в дерево
			if q.Fk_parent == v.Fk_parent {
				find = true
				trees[i].Data = append(trees[i].Data, subclass)
				break
			}
		}
		if !find {
			trees = append(trees, temp)
		}
	}
	tree := entity.FullTree{}
	tree.Value = "Все"
	tree.Id = 0
	tree.Open = true
	tree.Equipment = trees
	fullTree = append(fullTree, tree)
	return
}

//добавление оборудования
func (e *EquipmentModel) AddEquipment(DB *sql.DB, equipmentEntity entity.Equipment) (equipment entity.Equipment, err error) {
	var equipmentMapper mappers.Equipment
	equipmentEntity.StatusI = 0
	lastInsertedId, err := equipmentMapper.AddEquipment(DB, equipmentEntity)
	if err != nil {
		return
	}
	equipment, err = equipmentMapper.GetEquipmentById(DB, lastInsertedId)
	if err != nil {
		return
	}
	equipment.Status="На складе"
	equipment.UserFIO="Отсутствует"
	return
}

func (e *EquipmentModel) GetEquipmentByUser(DB *sql.DB, equipment entity.Equipment) (equipments []entity.Equipment, err error) {
	eqMapper := mappers.Equipment{}

	equipments, err = eqMapper.GetEquipmentsByUserId(DB, equipment)
	if err != nil {
		return
	}
	return
}

//удаление оборудования
func (e *EquipmentModel) DeleteEquipment(DB *sql.DB, equipment entity.Equipment) (result entity.Equipment, err error) {

	eqMapper := mappers.Equipment{}
	eventMapper := mappers.InventoryEvent{}
	_, err = eventMapper.DeleteEventByEmployee(DB, entity.Employee{Id: equipment.Id})
	if err != nil {
		return
	}
	deletedId, err := eqMapper.DeleteEquipment(DB, equipment)
	if err != nil {

		return
	}
	result.Id = deletedId
	return
}
