import {registerUserForm} from './registerUserForm.js';
import {Employee} from '../../../models/MEmployeeModel.js';
import {UpdateUserForm, UsersDatatable} from '../../const.js';
import {updateUserForm} from './updateUserForm.js';

export const UsersToolbar = {
    view: "toolbar",
    id: "UsersToolbar",
    cols: [
        {view: "button", id: "adduser", value: "Добавить сотрудника", width: 200, height: 50, click: addEmployee},
        {
            view: "button",
            id: "updateUser",
            value: "Изменить данные о сотруднике",
            width: 200,
            height: 50,
            click: updateEmployee
        },
        {view: "button", id: "deleteUser", value: "Удалить сотрудника", width: 200, height: 50, click: deleteEmoloyee},

        {view: "button", id: "resetPassword", value: "Сбросить пароль", width: 200, height: 50, click: resetPassword},
    ],
}

//добавления пользователя
function addEmployee() {
    registerUserForm.show();
}

//обновление данных пользователя
function updateEmployee() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        $$(UpdateUserForm).setValues({
            name: row.name,
            surname: row.surname,
            patronymic: row.patronymic,
            login: row.login,
        });
        updateUserForm.show();
    } else {
        webix.message("not selected item");
    }
}

//удаление пользователя
function deleteEmoloyee() {

    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        webix.confirm({
            title: "Удаление сотрудника",
            text: "Вы уверены?"
        }).then(() => {
            let user = new Employee();
            let promise = user.deleteUser(row);
            promise.then(response => {
                return response.json();
            }).then(result => {
                if (result.Error == "") {
                    $$(UsersDatatable).remove(result.Data.id);
                } else {
                    webix.message(result.Error);
                }
            })
        });
    } else {
        webix.message("not selected item");
    }
}

//сброс пароля
function resetPassword() {
    webix.confirm({
        title: "Сброс пароля",
        text: "Вы уверены?"
    }).then(() => {

        let row = $$(UsersDatatable).getSelectedItem();
        if (row) {
            let user = new Employee();
            console.log(user);
            let promise = user.resetPassword(row);
            promise.then(response => {
                return response.json();
            }).then(result => {
                if (result.Reject == null) {
                    webix.message("success reset");
                } else {
                    webix.message(result.Reject);
                }
            })
        } else {
            webix.message("not selected item");
        }
    });
}
