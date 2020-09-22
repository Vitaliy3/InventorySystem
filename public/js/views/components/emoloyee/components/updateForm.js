import {Employee} from '../../../../models/MEmployeeModel.js';
import {UpdateUserForm, UsersDatatable} from '../../../const.js';

//форма редактирования данных сотрудника
export const updateForm = webix.ui({
    view: "window",
    width: 400,
    height: 500,
    position:"center",
    head: "Редактирование данных сотрудника",
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
            {view: "text", name: "password",type:"password", label: "Пароль", labelWidth: 90},
            {
                margin: 5, cols: [
                    {view: "button", label: "Подтвердить", type: "form", click: updateEmployee},
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
        }
    }
});

//вызывается при подтвержедении обновления данных о сотруднике
function updateEmployee() {
    if ($$(UpdateUserForm).validate()) {
        let formValues = $$(UpdateUserForm).getValues();
        let row = $$(UsersDatatable).getSelectedItem();
        let user = new Employee();

        row.name = formValues.name;
        row.surname = formValues.surname;
        row.patronymic = formValues.patronymic;
        row.login = formValues.login;
        row.password = formValues.password;
        let promise = user.update(row);

        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                let datatable = $$(UsersDatatable);
                datatable.updateItem(row.id, result.Data);
                console.log(result.Data);
                updateForm.hide();
                webix.message("Успешное обновление");
            } else {
                webix.message(result.Error);
            }
        })
    }
}

//вызывается при отмене обновления данных о сотруднике
function cancel() {
    updateForm.hide();
}