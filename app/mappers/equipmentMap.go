package mappers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type EquipmentTable struct {
	Id              int           ` json:"id" `
	Fk_parent       int           ` json:"class" `
	Fk_class        int           ` json:"subclass" `
	Fk_user         sql.NullInt64 ` json:"user" `
	InventoryNumber string        ` json:"inventoryNumber" `
	EquipmentName   string        ` json:"name" `
	Status          int           ` json:"status" `
	Subclass        string
	Class           string
}

//получение одной еденицы оборудования
func (e *EquipmentTable) GetEquipmentById(DB *sql.DB, id int) (eq EquipmentTable, err error) {
	fmt.Println("getId", id)
	row := DB.QueryRow("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.id =$1", id)
	err = row.Scan(&eq.Id, &eq.Fk_class, &eq.Fk_user, &eq.InventoryNumber, &eq.EquipmentName, &eq.Status, &eq.Fk_parent)
	if err != nil {
		fmt.Println("errGetById:", err)
		return
	}
	return
}

//получение всего оборудования
func (e *EquipmentTable) GetAllEquipments(DB *sql.DB) (equipments []EquipmentTable, err error) {
	rows, err := DB.Query("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Id, &e.Fk_class, &e.Fk_user, &e.InventoryNumber, &e.EquipmentName, &e.Status, &e.Fk_parent)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			return
		}
		equipments = append(equipments, *e)
	}
	return
}

//все товары у сотрудника
func (e *EquipmentTable) GetEquipmentsByUser(DB *sql.DB, userId int) (equipments []EquipmentTable, err error) {

	rows, err := DB.Query("select equipments.id,fk_class,c2.id,fk_user,inventoryNumber,equipmentName,status from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.fk_user=$1", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Id, &e.Fk_class, &e.Fk_parent, &e.Fk_user, &e.InventoryNumber, &e.EquipmentName, &e.Status)
		if err != nil {
			return
			continue
		}
		equipments = append(equipments, *e)
	}
	return
}

//все товары на складе
func (e *EquipmentTable) GetEquipmentsInStore(DB *sql.DB) (equipments []EquipmentTable, err error) {
	rows, err := DB.Query("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.status =0")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Id, &e.Fk_class, &e.Fk_user, &e.InventoryNumber, &e.EquipmentName, &e.Status, &e.Fk_parent)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			return
		}
		equipments = append(equipments, *e)
	}
	return
}

//получение всех классов и подклассов
func (e *EquipmentTable) GetFullTree(DB *sql.DB) (equipments []EquipmentTable, err error) {
	rows, err := DB.Query("select c1.fk_parent ,c1 .id, classes.name,c1.name from classes join classes c1 on classes.id =c1.fk_parent")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Fk_parent, &e.Fk_class, &e.Class, &e.Subclass)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}
		equipments = append(equipments, *e)
	}
	return
}

//добавление оборудования
func (e *EquipmentTable) AddEquipment(DB *sql.DB) (lastInsertedId int, err error) {
	err = DB.QueryRow("insert into equipments (fk_class,inventoryNumber,equipmentName,status)values($1,$2,$3,$4) returning id",
		e.Fk_class, e.InventoryNumber, e.EquipmentName, e.Status).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

func (e *EquipmentTable) DragToUser(DB *sql.DB, fk_user int, id int) (lastInsertedId int, err error) {
	err = DB.QueryRow("update equipments set fk_user=$1,status=1 where id=$2 returning id", id, fk_user).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}
func (e *EquipmentTable) NewEvent(DB *sql.DB, fk_user sql.NullInt64, fk_equipment int, action string) (lastInsertedId int, err error) {
	err = DB.QueryRow("insert into inventoryEvents (fk_equipment,fk_user,actionevent,date) values($1,$2,$3,'now')", fk_equipment, fk_user, action).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

func (e *EquipmentTable) DragToStore(DB *sql.DB, id int) (lastInsertedId int, err error) {
	err = DB.QueryRow("update equipments set fk_user=null,status=0 where id=$1 returning id", id).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

//обновление данных об оборудовании
func (e *EquipmentTable) UpdateEquipment(DB *sql.DB) (lastUpdatedId int, err error) {
	err = DB.QueryRow("update equipments set equipmentname=$1, inventorynumber=$2 where id=$3 returning id", e.EquipmentName, e.InventoryNumber, e.Id).Scan(&lastUpdatedId)
	if err != nil {
		return
	}
	return
}

//удаление оборудования
func (e *EquipmentTable) DeleteEquipment(DB *sql.DB, id int) (rowsAffected int64, err error) {
	result, err := DB.Exec("delete from equipments where id=$1", id)
	if err != nil {
		fmt.Println("DeleteEquipmentMapper", err)
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

//списывание оборудования
func (e *EquipmentTable) WriteEquipment(DB *sql.DB, id int) (updatedElementId int, err error) {
	err = DB.QueryRow("update equipments set status=2 where id=$1 returning id", id).Scan(&updatedElementId)
	if err != nil {
		return
	}
	return
}
