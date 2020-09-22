package providers

import (
	"database/sql"
	"myapp/app/entity"
	"myapp/app/mappers"
)

type InventoryEvent struct {
	entity.InventoryEvent
	invEventMapper mappers.InventoryEvent
}

//получение всех событий
func (e *InventoryEvent) GetAll(DB *sql.DB) (events []entity.InventoryEvent, err error) {
	equipmentMapper := mappers.Equipment{}
	employeeMapper := mappers.Employee{}

	allEquip, err := equipmentMapper.GetAll(DB)
	if err != nil {
		return
	}

	allEmployees, err := employeeMapper.GetAll(DB)
	if err != nil {

		return
	}

	events, err = e.invEventMapper.GetAll(DB)
	if err != nil {

		return
	}

	for i, _ := range events {
		for _, m := range allEmployees {
			if int(events[i].Fk_user.Int64) == m.Id {
				events[i].UserFIO = m.Surname + " " + m.Surname + " " + m.Patronymic
			}
		}
		if events[i].UserFIO == "" {
			events[i].UserFIO = "Отсутствует"
		}
		for _, n := range allEquip {
			if events[i].Fk_equipment == n.Id {
				events[i].Equipment = n.EquipmentName
			}
		}
	}
	return
}

//получение событий за определенный промежуток времени
func (e *InventoryEvent) GetForDate(DB *sql.DB, dateStart, dateEnd string) (allEvents []entity.InventoryEvent, err error) {
	equipmentMapper := mappers.Equipment{}
	employeeMapper := mappers.Employee{}

	invEvent := InventoryEvent{}
	allEquip, err := equipmentMapper.GetAll(DB)
	if err != nil {
		return
	}

	allEmployees, err := employeeMapper.GetAll(DB)
	if err != nil {
		return
	}

	allEvents, err = e.invEventMapper.GetForDate(DB, dateStart, dateEnd)
	if err != nil {
		return
	}
	for i, _ := range allEvents {
		for _, m := range allEmployees {
			if int(allEvents[i].Fk_user.Int64) == m.Id {
				invEvent.UserFIO = m.Surname + " " + m.Surname + " " + m.Patronymic
			}
		}
		if allEvents[i].UserFIO == "" {
			allEvents[i].UserFIO = "Отсутствует"
		}
		for _, n := range allEquip {
			if allEvents[i].Fk_equipment == n.Id {
				allEvents[i].Equipment = n.EquipmentName
			}
		}
	}
	return
}
