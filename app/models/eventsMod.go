package models

import (
	"database/sql"
	"fmt"
	"myapp/app/entity"
	"myapp/app/mappers"
)

type InventoryEvent struct {
	entity.InventoryEvent
}

func (e *InventoryEvent) GetAllEvents(DB *sql.DB) (allEvents []entity.InventoryEvent, err error) {
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
	allEvents, err = inventoryEventMapper.GetAllEvents(DB)
	if err != nil {
		return
	}
	for i, _ := range allEvents {

		for _, m := range allEmployees {
			fmt.Println("here:", int(allEvents[i].Fk_user.Int64))
			fmt.Println("here1:", m.Id)

			if int(allEvents[i].Fk_user.Int64) == m.Id {
				allEvents[i].UserFIO = m.Surname + " " + m.Surname + " " + m.Patronymic
			}
		}
		for _, n := range allEquip {

			if allEvents[i].Fk_equipment == n.Id {
				allEvents[i].Equipment = n.EquipmentName
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
		for _, n := range allEquip {

			if allEvents[i].Fk_equipment == n.Id {
				invEvent.Equipment = n.EquipmentName
			}
		}
	}
	return
}
