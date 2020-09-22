import {registerUserForm} from "./register.js";
import {Employee} from '../../../../models/MEmployeeModel.js';
import {
    DeleteEmployeeButton,
    EmployeeToolbar,
    ResetPasswordButton,
    UpdateEmployeeButton,
    UpdateUserForm,
    UsersDatatable
} from '../../../const.js';
import {updateForm} from './updateForm.js';

//панель инструметов для отображения учета сотрудников
export const UsersToolbar = {
    view: "toolbar",
    id: EmployeeToolbar,
    cols: [
        {
            view: "button",
            id: "adduser",
            value: "Добавить сотрудника",
            width: 200,
            height: 50,
            click: ShowForm_AddEmployee
        },
        {
            view: "button",
            id: UpdateEmployeeButton,
            value: "Изменить данные о сотруднике",
            width: 200,
            height: 50,
            click: updateEmployee
        },
        {
            view: "button",
            id: DeleteEmployeeButton,
            value: "Удалить сотрудника",
            width: 200,
            height: 50,
            click: deleteEmployee
        },
        {
            view: "button",
            id: ResetPasswordButton,
            value: "Сбросить пароль",
            width: 200,
            height: 50,
            click: resetPassword
        },
    ],
};

//показывает форму для добавления пользователя
function ShowForm_AddEmployee() {
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
        updateForm.show();
    } else {
        webix.message("Пользователь не выбран");
    }
}

//удаление пользователя
function deleteEmployee() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        webix.confirm({
            title: "Удаление сотрудника",
            text: "Вы уверены?"
        }).then(() => {
            let user = new Employee();
            let promise = user.delete(row);
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
        webix.message("Пользователь не выбран");
    }
}

//сброс пароля
function resetPassword() {
    let row = $$(UsersDatatable).getSelectedItem();
    if (row) {
        webix.confirm({
            title: "Сброс пароля",
            text: "Вы уверены?"
        }).then(() => {
            let user = new Employee();
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
        });
    } else {
        webix.message("Пользователь не выбран");
    }
}
