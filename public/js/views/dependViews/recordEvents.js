import {UserEventsDatatable} from '../const.js';
import {EmployeeEvent} from "../../models/MEmployeeEvents.js";

export const recordEventsDatapicker = {
    view: "toolbar",
    scroll: false,
    width: 320,

    elements: [
        {view: "datepicker", label: "Начальная дата:", name: "start", stringResult: true, format: "%d  %M %Y",labelWidth:125,width:250},
        {view: "datepicker", label: "Конечная дата:", name: "end", stringResult: true, format: "%d  %M %Y",labelWidth:125,width:250},
        {view: "button", value: "Поиск", click: filterDate,width:150}
    ]

};


function filterDate() {
    let dateFromTo = JSON.stringify(this.getParentView().getValues());
    let eventModel = new EmployeeEvent();
    let promise = eventModel.getEventsForDate(dateFromTo);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            $$(UserEventsDatatable).clearAll();
            $$(UserEventsDatatable).parse(result.DataArray);
        } else {
            webix.message(result.Error);
        }
    })
}

export const recordEvents = {
    view: "datatable",
    id: UserEventsDatatable,
    editable: true,
    editaction: "custom",
    select: true,
    columns: [
        {id: "user", header: ["Сотрудник", {content: "selectFilter"}], fillspace: true,},
        {id: "event", header: ["Событие", {content: "selectFilter"}], fillspace: true},
        {id: "date", header: ["Дата", {content: "selectFilter"}], fillspace: true},
        {id: "equipment", header: ["Оборудование", {content: "selectFilter"}], fillspace: true},
    ]
}