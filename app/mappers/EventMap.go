package mappers

import (
	"database/sql"
	"fmt"
	"myapp/app/entity"
)

type InventoryEvent struct {
	entity.InventoryEvent
}

func (e *Equipment) NewEvent(DB *sql.DB, event entity.InventoryEvent) (lastInsertedId int, err error) {
	fmt.Println("fk_user1",event.Fk_userI)
	if event.Fk_userI == 0 {
		event.Fk_user.Valid = false
	} else {
		event.Fk_user.Int64 = int64(event.Fk_userI)
		event.Fk_user.Valid=true
	}
	fmt.Println("fk_userINmap:", event.Fk_user)
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

func (e *InventoryEvent) DeleteEventByEmployee(DB *sql.DB, employee entity.Employee) (deleteId int, err error) {
	err = DB.QueryRow("delete from inventoryevents where fk_equipment=$1 returning id", employee.Id).Scan(&deleteId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

func (e *InventoryEvent) DeleteEventByFkUser(DB *sql.DB, employee entity.Employee) (deletedId int, err error) {
	err = DB.QueryRow("delete from inventoryevents where fk_user=$1 returning id", employee.Id).Scan(&deletedId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return
}

func (e *InventoryEvent) GetAllEvents(DB *sql.DB) (inventoryEvent []entity.InventoryEvent, err error) {
	rows, err := DB.Query("select id, fk_user, fk_equipment,actionEvent,to_char(date, 'DD-MM-YYYY') from inventoryEvents")
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	var event entity.InventoryEvent
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&event.Id, &event.Fk_user, &event.Fk_equipment, &event.ActionEvent, &event.Date)
		if err != nil {
			return
		}
		inventoryEvent = append(inventoryEvent, event)

	}
	return
}
func (e *InventoryEvent) GetEventsForDate(DB *sql.DB, startDate string, endDate string) (inventoryEvent []entity.InventoryEvent, err error) {
	rows, err := DB.Query("select id,fk_user,fk_equipment,actionEvent,to_char(date, 'DD-MM-YYYY') from inventoryEvents where date between $1 and $2 ", startDate, endDate)
	if err != nil {
		return
	}
	defer rows.Close()
	var event entity.InventoryEvent
	for rows.Next() {
		err = rows.Scan(&event.Id, &event.Fk_user, &event.Fk_equipment, &event.ActionEvent, &event.Date)
		if err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
			return
		}
		inventoryEvent = append(inventoryEvent, event)
	}
	return
}
