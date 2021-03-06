import {Employee} from '../../../../models/MEmployeeModel.js';
import {RegisterUserForm, UsersDatatable} from "../../../const.js";

//форма для регистрации пользователя
export const registerUserForm = webix.ui({
    view: "window",
    width: 400,
    position: "center",
    height: 500,
    head: "Регистрация сотрудника",
    body: {
        view: "form",
        id: RegisterUserForm,
        scroll: false,
        elements: [
            {view: "text", name: "name", label: "Имя", labelWidth: 90},
            {view: "text", name: "surname", label: "Фамилия", labelWidth: 90},
            {view: "text", name: "patronymic", label: "Отчество", labelWidth: 90},
            {view: "text", name: "login", label: "Логин", labelWidth: 90},
            {view: "text", type: "password", name: "password", label: "Пароль", labelWidth: 90},
            {
                margin: 5, cols: [
                    {view: "button", label: "Зарегистрировать", type: "form", click: register},
                    {view: "button", label: "Отмена", click: cancel}
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
            password(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
        }
    }
});

//вызыватся при подтверджении регистарции пользователя
function register() {
    if ($$(RegisterUserForm).validate()) {
        let formValues = $$(RegisterUserForm).getValues();
        let user = new Employee();
        let promise = user.register(formValues);
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(UsersDatatable).add(result.Data);
                $$(UsersDatatable).refreshColumns();
                webix.message("Успешная регистрация");

                $$(RegisterUserForm).clear();
                $$(RegisterUserForm).clearValidation();
            } else {
                webix.message(result.Error);
            }
        })
    }
}

function cancel() {
    $$(RegisterUserForm).clearValidation();
    registerUserForm.hide();
}