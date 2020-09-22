import {EmployeeEventsDatatable} from '../../const.js';

//представление: события выдачи
export const recordEvent = {
    view: "datatable",
    id: EmployeeEventsDatatable,
    select: true,
    columns: [
        {id: "user", header: ["Сотрудник", {content: "selectFilter"}], fillspace: true,},
        {id: "event", header: ["Событие", {content: "selectFilter"}], fillspace: true},
        {id: "date", header: ["Дата", {content: "selectFilter"}], fillspace: true},
        {id: "equipment", header: ["Оборудование", {content: "selectFilter"}], fillspace: true},
    ]
}