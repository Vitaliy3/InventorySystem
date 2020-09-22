package controllers

import (
	"github.com/revel/revel"
	"myapp/app"
	"myapp/app/providers"
)

type Events struct {
	*revel.Controller
	eventModel      providers.InventoryEvent
	renderInterface app.RenderInterface
}

//получение всех событий
func (c Events) GetAll() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}

	result, err := c.eventModel.GetAll(app.Db)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result

	}
	return c.RenderJSON(c.renderInterface)
}

//получение всех событий за опредленный промежуток времени
func (c Events) GetForDate() revel.Result {
	if CheckPerm(c.Controller, "admin") {
	} else {
		return c.Render()
	}
	//начальная и конечная дата для выборки
	var dateStart = c.Params.Get("dateStart")
	var dateEnd = c.Params.Get("dateEnd")

	result, err := c.eventModel.GetForDate(app.Db, dateStart, dateEnd)
	if err != nil {
		c.renderInterface.Error = err.Error()
	} else {
		c.renderInterface.Data = result
	}
	return c.RenderJSON(c.renderInterface)
}
