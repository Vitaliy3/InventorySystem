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
            { view: "text", type: "password", name: "password", label: "Пароль", labelWidth: 90 },
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
            user(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
            status(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
            inventoryNumber(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
        }
    }
});
function registerUser() {
    promise.then(
        response => {
            $$("myList").add(response);
            webix.message("success add");
            $$('addItemForm').clear();
            $$('addItemForm').clearValidation();
        },
        err => {
            alert("err" + err);
        });
    let user = new user();

}
function closeForm() {
    registerUserForm.hide();
}