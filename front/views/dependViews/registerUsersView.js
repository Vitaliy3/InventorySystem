export const registerUsers = {
    view: "datatable",
    id: "usersList",
    editable: true,
    editaction: "custom",
    select: true,
    data: [{ user: "Ivan", surname: "Ivanovich", patronymic: "Ivanov", login: "IvanIvan" }],
    columns: [
        { id: "user", header: "Имя", width: 200, },
        { id: "surname", header: "Фамилия", width: 200 },
        { id: "patronymic", header: "Отчество", width: 100 },
        { id: "login", header: "Логин", width: 200 },
    ]
};