import { Employee } from '../../../models/MEmployeeModel.js';
export const authorizeForm = webix.ui({
    view: "window",
    width: 400,
    height: 500,
    head: "Регистрация сотрудника",
    autofit: false,
    body: {
        view: "form",
        id: "authorizeUserForm",
        scroll: false,
        width: 400,
        elements: [
            { view: "text", name: "login", label: "Логин", labelWidth: 90 },
            { view: "text", type: "password", name: "password", label: "Пароль", labelWidth: 90 },
            {
                margin: 5, cols: [
                    { view: "button", label: "Войти", type: "form", click: autohrize },
                    { view: "button", label: "Отмена", click: closeForm }
                ]
            }],
        rules: {
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
function autohrize() {
    if ($$("authorizeUserForm").validate()) {
        let formValues = $$("authorizeUserForm").getValues();
        let user = new Employee();
        let promise = user.authorize(formValues);
        promise.then(
            result => {
                webix.message("welcome " + result);
                $$('authorizeUserForm').clear();
                $$('authorizeUserForm').clearValidation();
                closeForm();
            }, err => {
                webix.message("not register:" + err);
            })
    }
}

function closeForm() {
    authorizeForm.hide();
}