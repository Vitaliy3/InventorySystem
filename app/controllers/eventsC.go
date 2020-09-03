package controllers

import (
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/models"
)

type Events struct {
	*revel.Controller
}

func (c Events) GetAllEvents() revel.Result {
	DataEvent := models.InventoryEvent{}
	renderInterface := app.RenderInterface{}
	result, err := DataEvent.GetAllEvents(app.DB,c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result

	}
	return c.RenderJSON(renderInterface)
}

func (c Events) GetEventsForDate() revel.Result {
	DataEvent := models.InventoryEvent{}
	renderInterface := app.RenderInterface{}

	result, err := DataEvent.GetEventsForDate(app.DB,c.Params)
	if err != nil {
		renderInterface.Error = err.Error()
	} else {
		renderInterface.Data = result

	}
	return c.RenderJSON(renderInterface)
}
