package mappers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"myapp/app/entity"
)

type Equipment struct {
	entity.Equipment
}

//получение одной еденицы оборудования
func (e *Equipment) GetById(DB *sql.DB, id int) (equipment entity.Equipment, err error) {
	row := DB.QueryRow("select distinct equipments.id,fk_class,c2.id,fk_user,inventorynumber,equipmentname,status,c2.name ,c1.name from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.id =$1", id)

	err = row.Scan(&equipment.Id, &equipment.Fk_class, &equipment.Fk_parent, &equipment.Fk_user, &equipment.InventoryNumber,
		&equipment.EquipmentName, &equipment.StatusI, &equipment.Class, &equipment.Subclass)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//получение всего оборудования
func (e *Equipment) GetAll(DB *sql.DB) (equipments []entity.Equipment, err error) {
	rows, err := DB.Query("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id")
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()

	var equipment entity.Equipment
	for rows.Next() {
		err = rows.Scan(&equipment.Id, &equipment.Fk_class, &equipment.Fk_user, &equipment.InventoryNumber, &equipment.EquipmentName,
			&equipment.StatusI, &equipment.Fk_parent)
		if err != nil {
			return
		}
		equipments = append(equipments, equipment)
	}
	return
}

//получение оборудования у сотрудника
func (e *Equipment) GetByUserId(DB *sql.DB, equipment entity.Equipment) (equipments []entity.Equipment, err error) {

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
			&equipment.EquipmentName, &equipment.Status, &equipment.Class, &equipment.Subclass)
		if err != nil {
			return
		}
		equipments = append(equipments, equipment)
	}
	return
}

//получение всего оборудования на складе
func (e *Equipment) GetInStore(DB *sql.DB) (equipments []entity.Equipment, err error) {
	rows, err := DB.Query("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.status =0")
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()
	var equipment entity.Equipment
	for rows.Next() {
		err = rows.Scan(&equipment.Id, &equipment.Fk_class, &equipment.Fk_user, &equipment.InventoryNumber, &equipment.EquipmentName,
			&equipment.Status, &equipment.Fk_parent)
		if err != nil {
			return
		}
		equipments = append(equipments, equipment)
	}
	return
}

//получение классов и подклассов для опреденного сотрудника
func (e *Equipment) GetTreeByUserId(DB *sql.DB, equipment entity.Equipment) (classes []entity.TreeClass, err error) {
	rows, err := DB.Query("select distinct c1.fk_parent ,c1.id,c2.name ,c1.name from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.fk_user=$1", equipment.Fk_userI)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()

	var class entity.TreeClass
	for rows.Next() {
		err = rows.Scan(&class.Fk_parent, &class.Fk_class, &class.ClassName, &class.Subclass)
		if err != nil {
		}
		classes = append(classes, class)
	}
	return
}

//получение всех классов и подклассов
func (e *Equipment) GetTree(DB *sql.DB) (classes []entity.TreeClass, err error) {
	rows, err := DB.Query("select c1.fk_parent ,c1 .id, classes.name,c1.name from classes join classes c1 on classes.id =c1.fk_parent")
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()

	var class entity.TreeClass
	for rows.Next() {
		err = rows.Scan(&class.Fk_parent, &class.Fk_class, &class.ClassName, &class.Subclass)
		if err != nil {
			return
		}
		classes = append(classes, class)
	}
	return
}

//добавление оборудования
func (e *Equipment) Add(DB *sql.DB, equipment entity.Equipment) (lastInsertedId int, err error) {
	fmt.Println("equipment",equipment)
	err = DB.QueryRow("insert into equipments (fk_class,inventoryNumber,equipmentName,status)values($1,$2,$3,$4) returning id",
		equipment.Fk_class, equipment.InventoryNumber, equipment.EquipmentName, equipment.StatusI).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

//выдача оборудования сотруднику
func (e *Equipment) DragToUser(DB *sql.DB, equipment entity.Equipment) (lastUpdatedId int, err error) {
	err = DB.QueryRow("update equipments set fk_user=$1,status=1 where id=$2 returning id", equipment.Fk_userI, equipment.Id).Scan(&lastUpdatedId)
	if err != nil {
		return
	}
	return
}

//перещемение оборудование на склад
func (e *Equipment) DragToStore(DB *sql.DB, equipment entity.Equipment) (lastUpdatedId int, err error) {
	err = DB.QueryRow("update equipments set fk_user=null,status=0 where id=$1 returning id", equipment.Id).Scan(&lastUpdatedId)
	if err != nil {
		return
	}
	return
}

//обновление данных об оборудовании
func (e *Equipment) Update(DB *sql.DB, equipment entity.Equipment) (lastUpdatedId int, err error) {
	err = DB.QueryRow("update equipments set equipmentname=$1, inventorynumber=$2 where id=$3 returning id",
		equipment.EquipmentName, equipment.InventoryNumber, equipment.Id).Scan(&lastUpdatedId)
	if err != nil {
		return
	}
	return
}

//удаление оборудования
func (e *Equipment) Delete(DB *sql.DB, equipment entity.Equipment) (deletedId int, err error) {
	err = DB.QueryRow("delete from equipments where id=$1 returning id", equipment.Id).Scan(&deletedId)
	if err != nil {
		return
	}
	return
}

//списывание оборудования
func (e *Equipment) Write(DB *sql.DB, equipment entity.Equipment) (updatedId int, err error) {
	err = DB.QueryRow("update equipments set status=2 where id=$1 returning id", equipment.Id).Scan(&updatedId)
	if err != nil {
		return
	}
	return
}
