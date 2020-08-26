import { User } from './../../../models/UserModel.js';
export const registerUserForm = webix.ui({
    view: "window",
    width: 400,
    height: 500,
    head: "Регистрация сотрудника",
    autofit: false,
    body: {
        view: "form",
        id: "registerUserForm",
        scroll: false,
        width: 400,
        elements: [
            { view: "text", name: "name", label: "Имя", labelWidth: 90 },
            { view: "text", name: "surname", label: "Фамилия", labelWidth: 90 },
            { view: "text", name: "patronymic", label: "Отчество", labelWidth: 90 },
            { view: "text", name: "login", label: "Логин", labelWidth: 90 },
            { view: "text", /*type: "password"*/ name: "password", label: "Пароль", labelWidth: 90 },
            {
                margin: 5, cols: [
                    { view: "button", label: "Зарегистрировать", type: "form", click: registerUser },
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
            password(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
        }
    }
});
function registerUser() {
    if ($$("registerUserForm").validate()) {
        let formValues = $$("registerUserForm").getValues();
        let user = new User(formValues);
        let promise = user.registerUser();
        promise.then(
            result => {
                console.log(result);
                $$("usersList").add(result);
                webix.message("success register");
                $$('registerUserForm').clear();
                $$('registerUserForm').clearValidation();
            }, err => {
                webix.message("not register:" + err);
            })
    }
}

function closeForm() {
    registerUserForm.hide();
}