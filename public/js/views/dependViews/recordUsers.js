import { UsersDatatable } from '../const.js';
export const usersList = {
    view: "datatable",
    id: UsersDatatable,
    editable: true,
    editaction: "custom",
    select: true,
    columns: [
        { id: "name", header: "Имя", width: 200, fillspace: true },
        { id: "surname", header: "Фамилия", width: 200, fillspace: true },
        { id: "patronymic", header: "Отчество", width: 100, fillspace: true },
        { id: "login", header: "Логин", width: 200, fillspace: true },
    ]
};