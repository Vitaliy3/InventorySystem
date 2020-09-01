package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"myapp/app"
)

type Events struct {
	*revel.Controller
}

func (c Events) GetAllEvents() revel.Result {
	DataEvent :=app.RenderDataEvents{}
	result, err := DataEvent.Data.GetAllEvents(c.Params)
	if err != nil {
		DataEvent.Error = err.Error()
	} else {
		DataEvent.DataArray = result

	}
	fmt.Println("DATA",DataEvent.DataArray)
	return c.RenderJSON(DataEvent)
}

func (c Events) GetEventsForDate() revel.Result {
	DataEvent :=app.RenderDataEvents{}

	result, err := DataEvent.Data.GetEventsForDate(c.Params)
	if err != nil {
		DataEvent.Error = err.Error()
	} else {
		fmt.Println(DataEvent.DataArray)
		DataEvent.DataArray = result

	}
	return c.RenderJSON(DataEvent)
}
