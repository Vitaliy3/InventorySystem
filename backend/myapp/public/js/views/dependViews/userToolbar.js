import { registerUserForm } from './forms/registerUserForm.js';
import { User } from '../../models/UserModel.js';
import { UsersDatatable } from '../const.js';
import { updateUserForm } from './forms/updateUserForm.js';
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

//добавления пользователя
function addUser() {
    registerUserForm.show({ x: 400, y: 150 });
}

//обновление данных пользователя
function updateUser() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        $$('updateUserForm').setValues({
            name: row.name,
            surname: row.surname,
            patronymic: row.patronymic,
            login: row.login
        });
        updateUserForm.show({ x: 400, y: 200 });
    } else {
        webix.message("not selected item");
    }
}

//удаление пользователя
function deleteUser() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        let user = new User();
        let promise = user.deleteUser(row);
        promise.then(
            result => {
                $$(UsersDatatable).remove(result);
            },
            err => {
                alert("err" + err);
            });
    } else {
        webix.message("not selected item");
    }
}

//сброс пароля
function resetPassword() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        let user = new User();
        console.log(user);
        let promise = user.resetPassword(row);
        promise.then(
            result => {
                webix.message("success reset");
            },
            err => {
                alert("err" + err);
            });
    } else {
        webix.message("not selected item");
    }
}
