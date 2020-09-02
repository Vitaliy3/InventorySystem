package models

import (
	"github.com/revel/revel"
	"myapp/app/mappers"
)

type InventoryEvent struct {
	Id        int    `json:"id"`
	UserFIO   string `json:"user"`
	Event     string `json:"event"`
	Date      string `json:"date"`
	Equipment string `json:"equipment"`
}

func (e *InventoryEvent) GetAllEvents(params *revel.Params) (allEvents []InventoryEvent, err error) {
	eventMapper := mappers.InventoryEvent{}
	equipmentMapper := mappers.EquipmentTable{}
	employeeMapper := mappers.Employee{}

	invEvent := InventoryEvent{}
	allEquip, err := equipmentMapper.GetAllEquipments()
	if err != nil {
		return
	}
	allEmployees, err := employeeMapper.GetAllEmployees()
	if err != nil {
		return
	}
	allEventsMap, err := eventMapper.GetAllEvents()
	if err != nil {
		return
	}

	for _, v := range allEventsMap {
		invEvent.Id = v.Id
		invEvent.Event = v.ActionEvent
		invEvent.Date = v.Date

		for _, m := range allEmployees {
			if v.Fk_user == m.Id {
				invEvent.UserFIO = m.Surname + " " + m.Surname + " " + m.Patronymic
			}
		}
		for _, n := range allEquip {

			if v.Fk_equipment == n.Id {
				invEvent.Equipment = n.EquipmentName
			}
		}
		allEvents = append(allEvents, invEvent)
	}
	return
}
func (e *InventoryEvent) GetEventsForDate(params *revel.Params) (allEvents []InventoryEvent, err error) {
	var dateStart = params.Get("dateStart")
	var dateEnd = params.Get("dateEnd")
	eventMapper := mappers.InventoryEvent{}
	equipmentMapper := mappers.EquipmentTable{}
	employeeMapper := mappers.Employee{}

	invEvent := InventoryEvent{}
	allEquip, err := equipmentMapper.GetAllEquipments()
	if err != nil {
		return
	}
	allEmployees, err := employeeMapper.GetAllEmployees()
	if err != nil {
		return
	}
	allEventsMap, err := eventMapper.GetEventsForDate(dateStart, dateEnd)
	if err != nil {
		return
	}
	for _, v := range allEventsMap {
		invEvent.Id = v.Id
		invEvent.Event = v.ActionEvent
		invEvent.Date = v.Date

		for _, m := range allEmployees {
			if v.Fk_user == m.Id {
				invEvent.UserFIO = m.Surname + " " + m.Surname + " " + m.Patronymic
			}
		}
		for _, n := range allEquip {

			if v.Fk_equipment == n.Id {
				invEvent.Equipment = n.EquipmentName
			}
		}
		allEvents = append(allEvents, invEvent)
	}
	return
}
