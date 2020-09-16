package mappers

import (
	"database/sql"
	_ "github.com/lib/pq"
	"myapp/app/entity"
)

type Equipment struct {
	entity.Equipment
}

//получение одной еденицы оборудования
func (e *Equipment) GetEquipmentById(DB *sql.DB, id int) (equipment entity.Equipment, err error) {
	row := DB.QueryRow("select distinct equipments.id,fk_class,c2.id,fk_user,inventorynumber,equipmentname,status,c2.name ,c1.name from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.id =$1", id)
	err = row.Scan(&equipment.Id, &equipment.Fk_class, &equipment.Fk_parent, &equipment.Fk_user, &equipment.InventoryNumber,
		&equipment.EquipmentName, &equipment.StatusI,&equipment.Class,&equipment.Subclass)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//получение всего оборудования
func (e *Equipment) GetAllEquipments(DB *sql.DB) (equipments []entity.Equipment, err error) {
	rows, err := DB.Query("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id")
	if err != nil {
		return
	}
	defer rows.Close()
	var equipment entity.Equipment
	for rows.Next() {
		err = rows.Scan(&equipment.Id, &equipment.Fk_class, &equipment.Fk_user, &equipment.InventoryNumber, &equipment.EquipmentName,
			&equipment.StatusI, &equipment.Fk_parent)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			return
		}
		equipments = append(equipments, equipment)
	}
	return
}

//все оборудование у сотрудника
func (e *Equipment) GetEquipmentsByUserId(DB *sql.DB, equipment entity.Equipment) (equipments []entity.Equipment, err error) {

	rows, err := DB.Query("select equipments.id,fk_class,c2.id,fk_user,inventorynumber,equipmentname,status,c2.name ,c1.name from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.fk_user=$1", equipment.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&equipment.Id, &equipment.Fk_class, &equipment.Fk_parent, &equipment.Fk_user, &equipment.InventoryNumber,
			&equipment.EquipmentName, &equipment.Status,&equipment.Class,&equipment.Subclass)
		if err != nil {
			return
			continue
		}
		equipments = append(equipments, equipment)
	}
	return
}

//все оборудование на складе
func (e *Equipment) GetEquipmentsInStore(DB *sql.DB) (equipments []entity.Equipment, err error) {
	rows, err := DB.Query("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.status =0")
	if err != nil {
		return
	}
	defer rows.Close()
	var equipment entity.Equipment
	for rows.Next() {
		err = rows.Scan(&equipment.Id, &equipment.Fk_class, &equipment.Fk_user, &equipment.InventoryNumber, &equipment.EquipmentName,
			&equipment.Status, &equipment.Fk_parent)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			return
		}
		equipments = append(equipments, equipment)
	}
	return
}

//получение классов и подклассов для опреденного сотрудника
func (e *Equipment) GetEmployeeTreeById(DB *sql.DB, equipment entity.Equipment) (equipments []entity.Equipment, err error) {
	rows, err := DB.Query("select distinct c1.fk_parent ,c1.id,c2.name ,c1.name from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.fk_user=$1", equipment.Fk_user1)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&equipment.Fk_parent, &equipment.Fk_class, &equipment.Class, &equipment.Subclass)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}
		equipments = append(equipments, equipment)
	}
	return
}

//получение всех классов и подклассов
func (e *Equipment) GetFullTree(DB *sql.DB) (equipments []entity.Equipment, err error) {
	rows, err := DB.Query("select c1.fk_parent ,c1 .id, classes.name,c1.name from classes join classes c1 on classes.id =c1.fk_parent")
	if err != nil {
		return
	}
	defer rows.Close()
	var equipment entity.Equipment
	for rows.Next() {
		err = rows.Scan(&equipment.Fk_parent, &equipment.Fk_class, &equipment.Class, &equipment.Subclass)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}
		equipments = append(equipments, equipment)
	}
	return
}

//добавление оборудования
func (e *Equipment) AddEquipment(DB *sql.DB, equipment entity.Equipment) (lastInsertedId int, err error) {
	err = DB.QueryRow("insert into equipments (fk_class,inventoryNumber,equipmentName,status)values($1,$2,$3,$4) returning id",
		equipment.Fk_class, equipment.InventoryNumber, equipment.EquipmentName, equipment.StatusI).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

//выдача оборудования сотруднику
func (e *Equipment) DragEquipmentToUser(DB *sql.DB, equipment entity.Equipment) (lastInsertedId int, err error) {
	err = DB.QueryRow("update equipments set fk_user=$1,status=1 where id=$2 returning id", equipment.Fk_user1, equipment.Id).Scan(&lastInsertedId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//перещемение оборудование на склад
func (e *Equipment) DragEquipmentToStore(DB *sql.DB, equipment entity.Equipment) (lastInsertedId int, err error) {
	err = DB.QueryRow("update equipments set fk_user=null,status=0 where id=$1 returning id", equipment.Id).Scan(&lastInsertedId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//обновление данных об оборудовании
func (e *Equipment) UpdateEquipment(DB *sql.DB, equipment entity.Equipment) (lastUpdatedId int, err error) {
	err = DB.QueryRow("update equipments set equipmentname=$1, inventorynumber=$2 where id=$3 returning id",
		equipment.EquipmentName, equipment.InventoryNumber, equipment.Id).Scan(&lastUpdatedId)
	if err != nil {
		return
	}
	return
}

//удаление оборудования
func (e *Equipment) DeleteEquipment(DB *sql.DB, equipment entity.Equipment) (deletedId int, err error) {
	err = DB.QueryRow("delete from equipments where id=$1 returning id", equipment.Id).Scan(&deletedId)
	if err != nil {
		return
	}
	return
}

//списывание оборудования
func (e *Equipment) WriteEquipment(DB *sql.DB, equipment entity.Equipment) (updatedElementId int, err error) {
	err = DB.QueryRow("update equipments set status=2 where id=$1 returning id", equipment.Id).Scan(&updatedElementId)
	if err != nil {
		return
	}
	return
}
