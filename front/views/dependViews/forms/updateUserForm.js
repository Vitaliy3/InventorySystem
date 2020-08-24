import { User } from './../../../models/UserModel.js';
import { UsersList } from './../../const.js';
export const updateUserForm = webix.ui({
    view: "window",
    width: 400,
    height: 500,
    head: "Регистрация сотрудника",
    autofit: false,
    body: {
        view: "form",
        id: "updateUserForm",
        scroll: false,
        width: 400,
        elements: [
            { view: "text", name: "name", label: "Имя", labelWidth: 90 },
            { view: "text", name: "surname", label: "Фамилия", labelWidth: 90 },
            { view: "text", name: "patronymic", label: "Отчество", labelWidth: 90 },
            { view: "text", name: "login", label: "Логин", labelWidth: 90 },
            {
                margin: 5, cols: [
                    { view: "button", label: "Зарегистрировать", type: "form", click: updateUser },
                    { view: "button", label: "Отмена", click: closeForm }
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
    if ($$("updateUserForm").validate()) {
        let formValues = $$("updateUserForm").getValues();
        let row = $$(UsersList).getSelectedItem();
        row.name = formValues.name;
        row.surname = formValues.surname;
        row.patronymic = formValues.patronymic;
        row.login = formValues.login;
        let user = new User(row);
        let promise = user.updateUser();
        promise.then(
            result => {
                let datatable = $$(UsersList);
                datatable.updateItem(result.id, result)
                updateUserForm.hide();
                webix.message("success update");
            }, err => {
                webix.message("not success update:" + err);
            })
    }
}

function closeForm() {
    updateUserForm.hide();
}