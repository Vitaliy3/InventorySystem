import {registerUserForm} from './registerUserForm.js';
import {Employee} from '../../../models/MUserModel.js';
import {UpdateUserForm, UsersDatatable} from '../../const.js';
import {updateUserForm} from './updateUserForm.js';

export const UsersToolbar = {
    view: "toolbar",
    id: "UsersToolbar",
    cols: [
        {view: "button", id: "adduser", value: "Добавить сотрудника", width: 200, height: 50, click: addUser},
        {
            view: "button",
            id: "updateUser",
            value: "Изменить данные о сотруднике",
            width: 200,
            height: 50,
            click: updateUser
        },
        {view: "button", id: "resetPassword", value: "Сбросить пароль", width: 200, height: 50, click: resetPassword},
    ],
}

//добавления пользователя
function addUser() {
    registerUserForm.show({x: 400, y: 150});
}

//обновление данных пользователя
function updateUser() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        $$(UpdateUserForm).setValues({
            name: row.name,
            surname: row.surname,
            patronymic: row.patronymic,
            login: row.login
        });
        updateUserForm.show({x: 400, y: 200});
    } else {
        webix.message("not selected item");
    }
}

//удаление пользователя
function deleteUser() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        let user = new Employee();
        let promise = user.deleteEmployee(row);
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(UsersDatatable).remove(result.Data.id);
            } else {
                webix.message(result.Error);
            }
        })
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
