import { registerUserForm } from './forms/registerUserForm.js';

export const UsersToolbar = {
    view: "toolbar",
    id: "UsersToolbar",
    cols: [
        { view: "button", id: "adduser", value: "Добавить сотрудника", width: 200, height: 50, click: addUser },
        { view: "button", id: "updateUser", value: "Изменить данные о сотруднике", width: 200, height: 50, click: updateUser },
        { view: "button", id: "deleteUser", value: "Удалить сотрудника", width: 200, height: 50, click: deleteUser },
        { view: "button", id: "resetPassword", value: "Сбросить пароль", width: 200, height: 50, click: resetPassword },
    ],
}
function addUser() {
    registerUserForm.show();
}
function updateUser() { }
function deleteUser() { }
function resetPassword() {
    let row = $$("usersList").getSelectedItem();
    if (row) {
        addItemForm.show({ x: 200, y: 200 });
    } else {
        webix.message("not selected item");
    }

}