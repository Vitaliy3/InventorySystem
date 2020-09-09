import {Employee} from '../../../models/MUserModel.js';
import {UpdateUserForm, UsersDatatable} from './../../const.js';

export const updateUserForm = webix.ui({
    view: "window",
    width: 400,
    height: 500,
    head: "Редактирование данных сотрудника",
    autofit: false,
    body: {
        view: "form",
        id: UpdateUserForm,
        scroll: false,
        width: 400,
        elements: [
            {view: "text", name: "name", label: "Имя", labelWidth: 90},
            {view: "text", name: "surname", label: "Фамилия", labelWidth: 90},
            {view: "text", name: "patronymic", label: "Отчество", labelWidth: 90},
            {view: "text", name: "login", label: "Логин", labelWidth: 90},
            {
                margin: 5, cols: [
                    {view: "button", label: "Подтвердить", type: "form", click: updateUser},
                    {view: "button", label: "Отмена", click: closeForm}
                ]
            }],
        rules: {
            name(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
            surname(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
            patronymic(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
            login(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
        }
    }
});

function updateUser() {
    if ($$(UpdateUserForm).validate()) {
        let formValues = $$(UpdateUserForm).getValues();
        let row = $$(UsersDatatable).getSelectedItem();
        row.name = formValues.name;
        row.surname = formValues.surname;
        row.patronymic = formValues.patronymic;
        row.login = formValues.login;
        let user = new Employee();
        let promise = user.updateUser(row);
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Reject == null) {
                let datatable = $$(UsersDatatable);
                datatable.updateItem(row.id, result.Data);
                updateUserForm.hide();
                webix.message("success update");
            } else {
                webix.message(result.Reject);
            }
        })
    }
}

function closeForm() {
    updateUserForm.hide();
}