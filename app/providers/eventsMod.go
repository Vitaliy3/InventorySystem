package providers

import (
	"database/sql"
	"myapp/app/entity"
	"myapp/app/mappers"
)

type InventoryEvent struct {
	entity.InventoryEvent
}

func (e *InventoryEvent) GetAllEvents(DB *sql.DB) (events []entity.InventoryEvent, err error) {
	inventoryEventMapper := mappers.InventoryEvent{}
	equipmentMapper := mappers.Equipment{}
	employeeMapper := mappers.Employee{}
	allEquip, err := equipmentMapper.GetAllEquipments(DB)
	if err != nil {
		return
	}
	allEmployees, err := employeeMapper.GetAllEmployees(DB)
	if err != nil {

		return
	}
	events, err = inventoryEventMapper.GetAllEvents(DB)
	if err != nil {

		return
	}
	for i, _ := range events {

		for _, m := range allEmployees {
			if int(events[i].Fk_user.Int64) == m.Id {
				events[i].UserFIO = m.Surname + " " + m.Surname + " " + m.Patronymic
			}
		}
		if events[i].UserFIO==""{
			events[i].UserFIO="Отсутствует"
		}
		for _, n := range allEquip {
			if events[i].Fk_equipment == n.Id {
				events[i].Equipment = n.EquipmentName
			}
		}
	}
	return
}

func (e *InventoryEvent) GetEventsForDate(DB *sql.DB, dateStart, dateEnd string) (allEvents []entity.InventoryEvent, err error) {

	eventMapper := mappers.InventoryEvent{}
	equipmentMapper := mappers.Equipment{}
	employeeMapper := mappers.Employee{}

	invEvent := InventoryEvent{}
	allEquip, err := equipmentMapper.GetAllEquipments(DB)
	if err != nil {
		return
	}
	allEmployees, err := employeeMapper.GetAllEmployees(DB)
	if err != nil {
		return
	}
	allEvents, err = eventMapper.GetEventsForDate(DB, dateStart, dateEnd)
	if err != nil {
		return
	}
	for i, _ := range allEvents {
		for _, m := range allEmployees {
			if int(allEvents[i].Fk_user.Int64) == m.Id {
				invEvent.UserFIO = m.Surname + " " + m.Surname + " " + m.Patronymic
			}
		}
		if allEvents[i].UserFIO==""{
			allEvents[i].UserFIO="Отсутствует"
		}

		for _, n := range allEquip {
			if allEvents[i].Fk_equipment == n.Id {
				allEvents[i].Equipment = n.EquipmentName
			}
		}
	}
	return
}
