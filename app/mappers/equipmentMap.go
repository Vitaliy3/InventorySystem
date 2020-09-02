package mappers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host   = "127.0.0.1"
	port   = 5433
	user   = "postgres"
	dbname = "dbtest"
)

type EquipmentTable struct {
	Id              int            ` json:"id" `
	Fk_parent       int            ` json:"class" `
	Fk_class        int            ` json:"subclass" `
	Fk_user         sql.NullString ` json:"user" `
	InventoryNumber string         ` json:"inventoryNumber" `
	EquipmentName   string         ` json:"name" `
	Status          int            ` json:"status" `
	Subclass        string
	Class           string
}

var db *sql.DB

func OpenConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("errOpen:", err)
	}
}

//получение одной еденицы оборудования
func (e *EquipmentTable) GetEquipmentById(id int) (eq EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	fmt.Println("getId", id)
	row := db.QueryRow("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.id =$1", id)
	err = row.Scan(&eq.Id, &eq.Fk_class, &eq.Fk_user, &eq.InventoryNumber, &eq.EquipmentName, &eq.Status, &eq.Fk_parent)
	if err != nil {
		fmt.Println("errGetById:", err)
		return
	}
	return
}

//получение всего оборудования
func (e *EquipmentTable) GetAllEquipments() (equipments []EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select equipments.id,fk_class,fk_user,inventoryNumber,equipmentName,status,c2.id from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.Id, &e.Fk_class, &e.Fk_user, &e.InventoryNumber, &e.EquipmentName, &e.Status, &e.Fk_parent)
		if err != nil {
			log.Println(err)
			continue
		}
		equipments = append(equipments, *e)
	}
	return
}

//все товары у сотрудника
func (e *EquipmentTable) GetEquipmentsByUser(userId int) (equipments []EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select equipments.id,fk_class,c2.id,fk_user,inventoryNumber,equipmentName,status from equipments join classes c1 on equipments.fk_class =c1.id join classes c2 on c1.fk_parent =c2.id where equipments.fk_user=$1", userId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.Id, &e.Fk_class, &e.Fk_parent, &e.Fk_user, &e.InventoryNumber, &e.EquipmentName, &e.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		equipments = append(equipments, *e)
	}
	return
}

//все товары на складе
func (e *EquipmentTable) GetEquipmentsInStore() (equipments []EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select id,fk_class,fk_user,inventoryNumber,equipmentName,status from equipments where status=0")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.Id, &e.Fk_class, &e.Fk_user, &e.InventoryNumber, &e.EquipmentName, &e.Status)
		if err != nil {
			log.Println(err)
			continue
		}
		equipments = append(equipments, *e)
	}
	return
}

//получение всех классов и подклассов
func (e *EquipmentTable) GetFullTree() (equipments []EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select c1.fk_parent ,c1 .id, classes.name,c1.name from classes join classes c1 on classes.id =c1.fk_parent")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.Fk_parent, &e.Fk_class, &e.Class, &e.Subclass)
		if err != nil {
			log.Println(err)
			continue
		}
		equipments = append(equipments, *e)
	}
	return
}

//добавление оборудования
func (e *EquipmentTable) AddEquipment() (lastInsertedId int, err error) {
	OpenConnection()
	defer db.Close()
	err = db.QueryRow("insert into equipments (fk_class,inventoryNumber,equipmentName,status)values($1,$2,$3,$4) returning id",
		e.Fk_class, e.InventoryNumber, e.EquipmentName, e.Status).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

func (e *EquipmentTable) DragToUser(fk_user int,id int) (lastInsertedId int, err error) {
	OpenConnection()
	defer db.Close()
	err = db.QueryRow("update equipments set fk_user=$1,status=1 where id=$2 returning id", fk_user, id).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

func (e *EquipmentTable) DragToStore(id int) (lastInsertedId int, err error) {
	OpenConnection()
	defer db.Close()
	err = db.QueryRow("update equipments set fk_user=null,status=0 where id=$1 returning id",id).Scan(&lastInsertedId)
	if err != nil {
		return
	}
	return
}

//обновление данных об оборудовании
func (e *EquipmentTable) UpdateEquipment() (lastUpdatedId int, err error) {
	OpenConnection()
	defer db.Close()
	err = db.QueryRow("update equipments set equipmentname=$1, inventorynumber=$2 where id=$3 returning id", e.EquipmentName, e.InventoryNumber, e.Id).Scan(&lastUpdatedId)
	if err != nil {
		return
	}
	return
}

//удаление оборудования
func (e *EquipmentTable) DeleteEquipment(id int) (rowsAffected int64, err error) {
	OpenConnection()
	defer db.Close()
	result, err := db.Exec("delete from equipments where id=$1", id)
	if err != nil {
		fmt.Println("DeleteEquipmentMapper", err)
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

//списывание оборудования
func (e *EquipmentTable) WriteEquipment(id int) (updatedElementId int, err error) {
	OpenConnection()
	defer db.Close()
	err = db.QueryRow("update equipments set status=2 where id=$1 returning id", id).Scan(&updatedElementId)
	if err != nil {
		return
	}
	return
}
