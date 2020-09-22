import {UserEvent} from "../../../../models/MEmployeeEvents.js";
import {EmployeeEventsDatatable} from "../../../const.js";

//панель инструментов с календарями и функией поиска по дате
export const recordEventsDatepicker = {
    view: "toolbar",
    width: 320,
    elements: [
        {
            view: "datepicker",
            label: "Начальная дата:",
            name: "start",
            stringResult: true,
            format: "%d  %M %Y",
            labelWidth: 125,
            width: 250
        },
        {
            view: "datepicker",
            label: "Конечная дата:",
            name: "end",
            stringResult: true,
            format: "%d  %M %Y",
            labelWidth: 125,
            width: 250
        },
        {view: "button", value: "Поиск", click: filterByDate, width: 150}
    ]
};

//получает начальную и конечную дату и отправляет запрос на сервер на получение оборудования
function filterByDate() {
    let dateFromTo = JSON.stringify(this.getParentView().getValues());
    let eventModel = new UserEvent();
    let promise = eventModel.getForDate(dateFromTo);

    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            $$(EmployeeEventsDatatable).clearAll();
            if (result.Data != null) {
                $$(EmployeeEventsDatatable).parse(result.Data);
            }
        } else {
            webix.message(result.Error);
        }
    });
}