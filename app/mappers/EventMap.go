package mappers

type InventoryEvent struct {
	Id           int
	Fk_user      int
	Fk_equipment int
	ActionEvent  string
	Date         string
}

func (e *InventoryEvent) GetAllEvents() (inventoryEvent []InventoryEvent, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select id, fk_user, fk_equipment,actionEvent,to_char(date, 'DD-MM-YYYY') from inventoryEvents")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Id, &e.Fk_user, &e.Fk_equipment, &e.ActionEvent, &e.Date)
		if err != nil {
			return
		}
		inventoryEvent = append(inventoryEvent, *e)

	}
	return
}
func (e *InventoryEvent) GetEventsForDate(startDate string, endDate string) (inventoryEvent []InventoryEvent, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select id,fk_user,fk_equipment,actionEvent,to_char(date, 'DD-MM-YYYY') "+
		"from inventoryEvents where date between $1 and $2 ", startDate, endDate)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.Id, &e.Fk_user, &e.Fk_equipment, &e.ActionEvent, &e.Date)
		if err != nil {
			return
		}
		inventoryEvent = append(inventoryEvent, *e)
	}
	return
}