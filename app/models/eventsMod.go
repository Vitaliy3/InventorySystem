package models

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app/mappers"
)

type InventoryEvent struct {
	Id        int
	UserFIO   string
	Event     string
	Equipment string
}

func (e *InventoryEvent) GetAllEvents(params *revel.Params) (events []mappers.InventoryEvent, err error) {
	eventMapper := mappers.InventoryEvent{}
	result, err := eventMapper.GetAllEvents()
	fmt.Println(result)

	if err != nil {
		return
	}
	events = result
	return
}
func (e *InventoryEvent) GetEventsForDate(params *revel.Params) (events []mappers.InventoryEvent, err error) {
	var dateStart = params.Get("dateStart")
	var dateEnd = params.Get("dateEnd")
	fmt.Println("dates",dateStart)
	fmt.Println("dates",dateEnd)
	eventMapper := mappers.InventoryEvent{}
	result, err := eventMapper.GetEventsForDate(dateStart, dateEnd)
	if err != nil {
		return
	}
	events = result
	return
}
