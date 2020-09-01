package mappers

type InventoryEvent struct {
	Id          int    `json:"id"`
	Username    string `json:"user"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Equipment   string `json:"equipment"`
	ActionEvent string `json:"event"`
	Date        string `json:"date"`
}

func (e *InventoryEvent) GetAllEvents() (inventoryEvent []InventoryEvent, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select i.id,u.username,u.surname ,u.patronymic,e.equipmentName,i.actionEvent,to_char(i.date, 'DD-MM-YYYY') " +
		"from inventoryEvents i join users u on i.fk_user =u.id join equipments e on i.fk_equipment =e.id")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.Id, &e.Username, &e.Surname, &e.Patronymic, &e.Equipment, &e.ActionEvent, &e.Date)
		if err != nil {

		}
		inventoryEvent = append(inventoryEvent, *e)

	}
	return
}
func (e *InventoryEvent) GetEventsForDate(startDate string, endDate string) (inventoryEvent []InventoryEvent, err error) {
	OpenConnection()
	defer db.Close()
	rows, err := db.Query("select i.id,u.username,u.surname ,u.patronymic,e.equipmentName,i.actionEvent,to_char(i.date, 'DD-MM-YYYY') "+
		"from inventoryEvents i join users u on i.fk_user =u.id join equipments e on i.fk_equipment =e.id where date between $1 and $2 ", startDate, endDate)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&e.Id, &e.Username, &e.Surname, &e.Patronymic, &e.Equipment, &e.ActionEvent, &e.Date)
		if err != nil {
			continue
		}
		inventoryEvent = append(inventoryEvent, *e)
	}
	return
}
