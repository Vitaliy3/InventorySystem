package providers

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
func (e *EquipmentModel) DragToUser(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	equipmentMapper := mappers.Equipment{}
	updatedRowId, err := equipmentMapper.DragEquipmentToUser(DB, equipmentIn)
	if err != nil {
		return
	}
	equipmentOut, err = equipmentMapper.GetEquipmentById(DB, updatedRowId)
	if err != nil {
		return
	}
	_, err = equipmentMapper.NewEvent(DB, entity.InventoryEvent{Fk_userI: equipmentIn.Fk_user1, Fk_equipment: equipmentIn.Id, ActionEvent: "Выдача сотруднику"})
	if err != nil {
		return
	}
	equipmentOut.Status = getStatus(equipmentIn.StatusI)
	return
}

//перемещение оборудования от сотрудника на склад
func (e *EquipmentModel) DragToStore(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	equipmentMapper := mappers.Equipment{}

	updatedRowId, err := equipmentMapper.DragEquipmentToStore(DB, equipmentIn)
	if err != nil {
		return
	}
	equipmentOut, err = equipmentMapper.GetEquipmentById(DB, updatedRowId)
	if err != nil {
		return
	}
	_, err = equipmentMapper.NewEvent(DB, entity.InventoryEvent{Fk_userI:0, Fk_equipment: equipmentOut.Id, ActionEvent: "Перемещение на склад"})
	if err != nil {
		return
	}
	equipmentOut.Status = getStatus(equipmentOut.StatusI)
	return
}

//изменение оборудования
func (e *EquipmentModel) UpdateEquipment(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	eqMapper := mappers.Equipment{}
	employee := mappers.Employee{}

	lastInsertedId, err := eqMapper.UpdateEquipment(DB, equipmentIn)
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
	equipmentOut.Id = row.Id
	equipmentOut.Fk_class = row.Fk_parent
	if user.Name != "" {
		equipmentOut.UserFIO = user.Name + " " + user.Surname + " " + user.Patronymic
	} else {
		equipmentOut.UserFIO = "Отсутсвует"
	}
	equipmentOut.Fk_class = row.Fk_class
	equipmentOut.EquipmentName = row.EquipmentName
	equipmentOut.InventoryNumber = row.InventoryNumber
	equipmentOut.Status = getStatus(row.StatusI)
	return
}

//списать оборудование
func (e *EquipmentModel) WriteEquipment(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	eqMapper := mappers.Equipment{}

	updatedId, err := eqMapper.WriteEquipment(DB, equipmentIn)
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

//func (e *EquipmentModel) GetEmployeeByUserId(DB *sql.DB, equipmentIn entity.Equipment) (equipments []entity.Equipment, err error) {
//	equipmentMapper := mappers.Equipment{}
//	employeeMapper := mappers.Employee{}
//	equipments, err = equipmentMapper.GetEquipmentsByUserId(DB, equipmentIn)
//	employees, err := employeeMapper.GetAllEmployees(DB)
//	for i, _ := range equipments {
//		equipments[i].Status = getStatus(equipments[i].StatusI)
//		for _, m := range employees {
//			if int(equipments[i].Fk_user.Int64) == m.Id {
//				equipments[i].UserFIO = m.Name + " " + m.Surname + " " + m.Patronymic
//			}
//		}
//		if equipments[i].UserFIO == "" {
//			equipments[i].UserFIO = "Отсутствует"
//		}
//	}
//	return
//}

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
func (e *EquipmentModel) GetEmployeeTreeByUserId(DB *sql.DB, equipmentIn entity.Equipment) (fullTree []entity.FullTree, err error) {
	equipmentMapper := mappers.Equipment{}
	var employeeTree []entity.TreeClass
	employeeTree, err = equipmentMapper.GetEmployeeTreeByUserId(DB, equipmentIn)
	if err != nil {
		return
	}
	var trees []entity.TreeClass
	for _, v := range employeeTree {
		var temp entity.TreeClass
		temp.Class = v.Class
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
	tree.Tree = trees
	fullTree = append(fullTree, tree)
	return
}

func (e *EquipmentModel) GetFullTree(DB *sql.DB) (fullTree []entity.FullTree, err error) {
	equipmentMapper := mappers.Equipment{}
	var treeClasses []entity.TreeClass
	treeClasses, err = equipmentMapper.GetFullTree(DB)
	if err != nil {
		return
	}
	var trees []entity.TreeClass
	for _, v := range treeClasses {
		var temp entity.TreeClass
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
	tree.Tree = trees
	fullTree = append(fullTree, tree)
	return
}

//добавление оборудования
func (e *EquipmentModel) AddEquipment(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	var equipmentMapper mappers.Equipment
	equipmentIn.StatusI = 0
	lastInsertedId, err := equipmentMapper.AddEquipment(DB, equipmentIn)
	if err != nil {
		return
	}
	equipmentOut, err = equipmentMapper.GetEquipmentById(DB, lastInsertedId)
	if err != nil {
		return
	}
	equipmentOut.Status="На складе"
	equipmentOut.UserFIO="Отсутствует"
	return
}

func (e *EquipmentModel) GetEquipmentByUserId(DB *sql.DB, equipmentIn entity.Equipment) (equipments []entity.Equipment, err error) {
	eqMapper := mappers.Equipment{}

	equipments, err = eqMapper.GetEquipmentsByUserId(DB, equipmentIn)
	if err != nil {
		return
	}
	return
}

//удаление оборудования
func (e *EquipmentModel) DeleteEquipment(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {

	eqMapper := mappers.Equipment{}
	eventMapper := mappers.InventoryEvent{}
	_, err = eventMapper.DeleteEventByEmployee(DB, entity.Employee{Id: equipmentIn.Id})
	if err != nil {
		return
	}
	deletedId, err := eqMapper.DeleteEquipment(DB, equipmentIn)
	if err != nil {

		return
	}
	equipmentOut.Id = deletedId
	return
}
