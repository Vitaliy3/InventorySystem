import { UserEvent, EventToolbar } from './const.js';
import { InventoryEvent } from '../models/UserEvent.js';

export const UserEventToobar = {
    view: "toolbar",
    id: EventToolbar,
    cols: [
        { view: "text", id: "dateFrom", label: "Начальная дата", labelwidth: 200, },
        { view: "text", id: "dateTo", label: "Конечная дата", labelwidth: 200, },
        { view: "button", value: "Найти", click: filterDate, width: 100 },
    ],
}
function filterDate() {
    let dateFrom = $$("dateFrom").getValue();
    let dateTo = $$("dateTo").getValue();
    let events = new InventoryEvent({ dateFrom: dateFrom, dateTo: dateTo });
    let promise = events.getEventsDate();
    promise.then(
        result => {
            $$(UserEvent).clearAll();
            $$(UserEvent).parse(result);
        }, err => { webix.message(err); }
    )
}

export const userEvents = {
    view: "datatable",
    id: UserEvent,
    editable: true,
    editaction: "custom",
    select: true,
    columns: [
        { id: "user", header: ["Сотрудник", { content: "selectFilter" }], fillspace: true, },
        { id: "event", header: ["событие", { content: "selectFilter" }], fillspace: true },
        { id: "product", header: ["Оборудование", { content: "selectFilter" }], fillspace: true },
    ]
}