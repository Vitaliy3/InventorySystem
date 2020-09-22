package mappers

import (
	"database/sql"
	"myapp/app/entity"
)

type InventoryEvent struct {
	entity.InventoryEvent
}

//создание события
func (e *Equipment) Create(DB *sql.DB, event entity.InventoryEvent) (lastInsertedId int, err error) {
	if event.Fk_userI == 0 {
		event.Fk_user.Valid = false
	} else {
		event.Fk_user.Int64 = int64(event.Fk_userI)
		event.Fk_user.Valid=true
	}

	err = DB.QueryRow("insert into inventoryEvents (fk_equipment,fk_user,actionevent,date) values($1,$2,$3,'now')",
		event.Fk_equipment, event.Fk_user, event.ActionEvent).Scan(&lastInsertedId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//удаление всех событий на определенное оборудование
func (e *InventoryEvent) DeleteByEquipmentId(DB *sql.DB, employee entity.Employee) (deleteId int, err error) {
	err = DB.QueryRow("delete from inventoryevents where fk_equipment=$1 returning id", employee.Id).Scan(&deleteId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//удаление всех событий привязанных к сотруднику
func (e *InventoryEvent) DeleteByFkUser(DB *sql.DB, employee entity.Employee) (deletedId int, err error) {
	err = DB.QueryRow("delete from inventoryevents where fk_user=$1 returning id", employee.Id).Scan(&deletedId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

//получение всех событий
func (e *InventoryEvent) GetAll(DB *sql.DB) (inventoryEvent []entity.InventoryEvent, err error) {
	rows, err := DB.Query("select id, fk_user, fk_equipment,actionEvent,to_char(date, 'DD-MM-YYYY') from inventoryEvents")
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()

	var event entity.InventoryEvent
	for rows.Next() {
		err = rows.Scan(&event.Id, &event.Fk_user, &event.Fk_equipment, &event.ActionEvent, &event.Date)
		if err != nil {
			return
		}
		inventoryEvent = append(inventoryEvent, event)
	}
	return
}

//получение событий за определенный промежуток времени
func (e *InventoryEvent) GetForDate(DB *sql.DB, startDate string, endDate string) (inventoryEvent []entity.InventoryEvent, err error) {
	rows, err := DB.Query("select id,fk_user,fk_equipment,actionEvent,to_char(date, 'DD-MM-YYYY') from inventoryEvents where date between $1 and $2 ", startDate, endDate)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	defer rows.Close()

	var event entity.InventoryEvent
	for rows.Next() {
		err = rows.Scan(&event.Id, &event.Fk_user, &event.Fk_equipment, &event.ActionEvent, &event.Date)
		if err != nil {
			return
		}
		inventoryEvent = append(inventoryEvent, event)
	}
	return
}
