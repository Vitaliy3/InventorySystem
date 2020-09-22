package providers

import (
	"database/sql"
	"myapp/app/entity"
	"myapp/app/mappers"
)

type EquipmentModel struct {
	entity.Equipment
	equipmentMapper mappers.Equipment
}
type FullTree struct {
	entity.FullTree
}

//перемещение оборудования со склада к пользователю
func (e *EquipmentModel) DragToUser(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	updatedRowId, err := e.equipmentMapper.DragToUser(DB, equipmentIn)
	if err != nil {

		return
	}

	equipmentOut, err = e.equipmentMapper.GetById(DB, updatedRowId)
	if err != nil {
		return
	}

	_, err = e.equipmentMapper.Create(DB, entity.InventoryEvent{Fk_userI: equipmentIn.Fk_userI, Fk_equipment: equipmentIn.Id, ActionEvent: "Выдача сотруднику"})
	if err != nil {

		return
	}

	equipmentOut.Status = getStatus(equipmentIn.StatusI)
	return
}

//перемещение оборудования от сотрудника на склад
func (e *EquipmentModel) DragToStore(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	updatedRowId, err := e.equipmentMapper.DragToStore(DB, equipmentIn)
	if err != nil {
		return
	}

	equipmentOut, err = e.equipmentMapper.GetById(DB, updatedRowId)
	if err != nil {
		return
	}

	_, err = e.equipmentMapper.Create(DB, entity.InventoryEvent{Fk_userI: 0, Fk_equipment: equipmentOut.Id, ActionEvent: "Перемещение на склад"})
	if err != nil {
		return
	}

	equipmentOut.Status = getStatus(equipmentOut.StatusI)
	return
}

//редактирование оборудования
func (e *EquipmentModel) Update(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	employee := mappers.Employee{}

	lastInsertedId, err := e.equipmentMapper.Update(DB, equipmentIn)
	if err != nil {
		return
	}

	row, err := e.equipmentMapper.GetById(DB, lastInsertedId)
	if err != nil {
		return
	}

	user, err := employee.GetById(DB, int(row.Fk_user.Int64))
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

//списание оборудованиея
func (e *EquipmentModel) Write(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	updatedId, err := e.equipmentMapper.Write(DB, equipmentIn)
	if err != nil {

		return
	}

	equipmentOut, err = e.equipmentMapper.GetById(DB, updatedId)
	if err != nil {

		return
	}

	if equipmentOut.StatusI == 2 {
		equipmentOut.Status = "Списано"
	}
	return
}

//получение оборудования, которое находится на складе
func (e *EquipmentModel) GetInStore(DB *sql.DB) (equipments []entity.Equipment, err error) {
	equipments, err = e.equipmentMapper.GetInStore(DB)
	for i, _ := range equipments {
		equipments[i].Status = getStatus(equipments[i].StatusI)
	}
	return
}

//получение всего оборудования
func (e *EquipmentModel) GetAll(DB *sql.DB) (equipments []entity.Equipment, err error) {
	employeeMapper := mappers.Employee{}
	equipments, err = e.equipmentMapper.GetAll(DB)
	employees, err := employeeMapper.GetAll(DB)

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

//получение древовидной структуры для пользователя
func (e *EquipmentModel) GetTreeByUserId(DB *sql.DB, equipmentIn entity.Equipment) (fullTree []entity.FullTree, err error) {
	var (
		treeClasses []entity.TreeClass
		trees       []entity.TreeClass
	)

	treeClasses, err = e.equipmentMapper.GetTreeByUserId(DB, equipmentIn)
	if err != nil {
		return
	}
	for _, v := range treeClasses {
		var temp entity.TreeClass
		temp.Fk_parent = v.Fk_parent
		temp.Open = true
		temp.ClassName = v.ClassName
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
	tree := entity.FullTree{
		Value: "Все",
		Id:    0,
		Open:  true,
		Tree:  trees,
	}
	fullTree = append(fullTree, tree)
	return
}

//получение древовидной структуры
func (e *EquipmentModel) GetTree(DB *sql.DB) (fullTree []entity.FullTree, err error) {
	var (
		treeClasses []entity.TreeClass
		trees       []entity.TreeClass
	)

	treeClasses, err = e.equipmentMapper.GetTree(DB)
	if err != nil {
		return
	}
	for _, v := range treeClasses {
		var temp entity.TreeClass
		temp.ClassName = v.ClassName
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
	tree := entity.FullTree{
		Value: "Все",
		Id:    0,
		Open:  true,
		Tree:  trees,
	}
	fullTree = append(fullTree, tree)
	return
}

//добавление оборудования
func (e *EquipmentModel) Add(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	equipmentIn.StatusI = 0

	lastInsertedId, err := e.equipmentMapper.Add(DB, equipmentIn)
	if err != nil {
		return
	}

	equipmentOut, err = e.equipmentMapper.GetById(DB, lastInsertedId)
	if err != nil {
		return
	}

	equipmentOut.Status = "На складе"
	equipmentOut.UserFIO = "Отсутствует"
	return
}

//получение оборудования у сотрудника
func (e *EquipmentModel) GetByUserId(DB *sql.DB, equipmentIn entity.Equipment) (equipments []entity.Equipment, err error) {
	eqMapper := mappers.Equipment{}

	equipments, err = eqMapper.GetByUserId(DB, equipmentIn)
	if err != nil {
		return
	}
	return
}

//удаление оборудования
func (e *EquipmentModel) Delete(DB *sql.DB, equipmentIn entity.Equipment) (equipmentOut entity.Equipment, err error) {
	eventMapper := mappers.InventoryEvent{}

	_, err = eventMapper.DeleteByEquipmentId(DB, entity.Employee{Id: equipmentIn.Id})
	if err != nil {
		return
	}

	deletedId, err := e.equipmentMapper.Delete(DB, equipmentIn)
	if err != nil {

		return
	}

	equipmentOut.Id = deletedId
	return
}
