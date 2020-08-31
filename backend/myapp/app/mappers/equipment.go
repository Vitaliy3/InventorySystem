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
		err := rows.Scan(&e.Id, &e.Fk_class, &e.Fk_user, &e.InventoryNumber, &e.EquipmentName, &e.Status,&e.Fk_parent)
		if err != nil {
			log.Println(err)
			continue
		}
		equipments = append(equipments, *e)
	}
	return
}

//все товары у сотрудника
func (e *EquipmentTable) GetAUserEquipments() (equipments []EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select * from equipments where fk_user=?", e.Id)
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

//все товары на складе
func (e *EquipmentTable) GetStoreEquipments() (equipments []EquipmentTable, err error) {
	rows, err := db.Query("select * from equipments where fk_class=NULL")
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

//обновление данных об оборудовании
func (e *EquipmentTable) UpdateEquipment() (rowsAffected int64, err error) {
OpenConnection()
defer db.Close()
	rows, err := db.Exec("update equipments set equipmentname=$1, inventorynumber=$2 where id=$3",e.EquipmentName,e.InventoryNumber,e.Id)
	if err != nil {
		fmt.Println("err in updateEqMapper",err)
		return
	}
	rowsAffected, err = rows.RowsAffected()
	return
}

//удаление оборудования
func (e *EquipmentTable) DeleteEquipment(id int) (rowsAffected int64, err error) {
	OpenConnection()
	defer db.Close()
	result, err := db.Exec("delete from equipments where id=$1", id)
	if err != nil {
		fmt.Println("DeleteEquipmentMapper",err)
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}



//списывание оборудования
func (e *EquipmentTable) WriteEquipment(id int) (rowsAffected int64, err error) {
	OpenConnection()
	defer db.Close()
	result, err := db.Exec("update equipments set status=2 where id=$1", id)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}
