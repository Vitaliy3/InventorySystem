package mappers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "127.0.0.1"
	port     = 5433
	user     = "postgres"
	password = ""
	dbname   = "mydb"
)

type EquipmentTable struct {
	Id              int
	Fk_class        int
	Fk_user         sql.NullString
	InventoryNumber string
	EquipmentName   string
	Status          int
}

var db *sql.DB

func OpenConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
}

//получение одной еденицы оборудования
func (e *EquipmentTable) GetEquipmentById(id int) (eq EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	row := db.QueryRow("select * from equipments where id=$1", id)
	if err != nil {
		return
	}
	err = row.Scan(&eq.Id, &eq.Fk_class, &eq.Fk_user, &eq.InventoryNumber, &eq.EquipmentName, &eq.Status)
	if err != nil {
		return
	}
	return
}

//получение всего оборудования
func (e *EquipmentTable) GetAllEquipments() (equipments []EquipmentTable, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select * from equipments")
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

//добавление оборудования
func (e *EquipmentTable) AddEquipment() (rowsAffected int64, err error) {
	rows, err := db.Exec("insert into equipments (fk_class,inventoryNumber,equipmentName,status)values("+
		"'$1','$2','$3',$4)", e.Fk_class, e.InventoryNumber, e.EquipmentName, e.Status)
	if err != nil {
		return
	}
	rowsAffected, err = rows.RowsAffected()
	return
}

//обновление данных об оборудовании
func (e *EquipmentTable) UpdateEquipment(id int, equipmentname string, inventoryNumber string) (rowsAffected int64, err error) {

	rows, err := db.Exec("update equipments set equipmentname=$1, inventorynumber=$2 where id=$3", equipmentname, inventoryNumber, id)
	if err != nil {
		return
	}
	rowsAffected, err = rows.RowsAffected()
	return
}

//удаление оборудования
func (e *EquipmentTable) DeleteEquipment() (rowsAffected int64, err error) {
	result, err := db.Exec("delete from equipments where id=$1", e.Id)
	if err != nil {
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
