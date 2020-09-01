import {Employee} from '../../../models/MUserModel.js';

export const registerEmployeeForm = webix.ui({
    view: "window",
    width: 400,
    height: 500,
    head: "Регистрация сотрудника",
    autofit: false,
    body: {
        view: "form",
        id: "registerEmployeeForm",
        scroll: false,
        width: 400,
        elements: [
            {view: "text", name: "name", label: "Имя", labelWidth: 90},
            {view: "text", name: "surname", label: "Фамилия", labelWidth: 90},
            {view: "text", name: "patronymic", label: "Отчество", labelWidth: 90},
            {view: "text", name: "login", label: "Логин", labelWidth: 90},
            {view: "text", /*type: "password"*/ name: "password", label: "Пароль", labelWidth: 90},
            {
                margin: 5, cols: [
                    {view: "button", label: "Зарегистрировать", type: "form", click: registerUser},
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
            password(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
        }
    }
});

function registerUser() {
    if ($$("registerEmployeeForm").validate()) {
        let formValues = $$("registerEmployeeForm").getValues();
        let user = new Employee();
        let promise = user.registerUser(formValues);
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$("usersList").add(result.Data);
                $$("usersList").refreshColumns();
                webix.message("success register");
                $$('registerEmployeeForm').clear();
                $$('registerEmployeeForm').clearValidation();
            } else {
                webix.message(result.Error);
            }
        })
    }
}

function closeForm() {
    registerEmployeeForm.hide();
}