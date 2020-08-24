import { UserEvent } from './const.js';
export const userEvents = {
    view: "datatable",
    id: UserEvent,
    editable: true,
    editaction: "custom",
    select: true,

    columns: [
        { id: "user", header: ["Сотрудник", { content: "selectFilter" }], fillspace: true, },
        { id: "date", header: ["Deadline", { content: "dateRangeFilter" }], format: webix.i18n.dateFormatStr, auttowidth: true, fillspace: true },
        { id: "event", header: ["событие", { content: "selectFilter" }], fillspace: true },
        { id: "product", header: ["Оборудование", { content: "selectFilter" }], fillspace: true },
    ]
}