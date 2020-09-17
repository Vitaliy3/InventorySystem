package controllers

import (
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/providers"
)

type Events struct {
	*revel.Controller
}

func (c Events) GetAllEvents() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	DataEvent := providers.InventoryEvent{}
	renderInterface := app.RenderInterface{}
	result, err := DataEvent.GetAllEvents(app.DB)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result

	}
	return c.RenderJSON(renderInterface)
}

func (c Events) GetEventsForDate() revel.Result {
	DataEvent := providers.InventoryEvent{}
	renderInterface := app.RenderInterface{}
	var dateStart = c.Params.Get("dateStart")
	var dateEnd = c.Params.Get("dateEnd")
	result, err := DataEvent.GetEventsForDate(app.DB,dateStart,dateEnd)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result

	}
	return c.RenderJSON(renderInterface)
}
