import {Equipment} from '../../../../models/MEquipmentM.js';

//форма для редактирования оборудования
export const updateEquipmentForm = webix.ui({
    view: "window",
    width: 600,
    height: 500,
    position: "center",
    head: "Редактирование оборудования",
    body: {
        view: "form",
        id: "updateEquipmentForm",
        scroll: false,
        width: 600,
        elements: [
            {id: "name", view: "text", name: "name", label: "Название"},
            {id: "number", view: "text", name: "inventoryNumber", label: "Инвентарный номер"},

            {
                margin: 5, cols: [
                    {view: "button", label: "Подтвердить", type: "form", click: confirmUpdateItem},
                    {view: "button", label: "Отмена", click: cancel}
                ]
            }],
        rules: {
            name(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
            inventoryNumber(value) {
                if (webix.rules.isNotEmpty(value)) {
                    return true;
                }
            },
        },
    },

});

//вызывается при подтверждении на изменение данных об оборудовании, считывает данные из формы и отправляет запрос на сервер на обновление данных
function confirmUpdateItem() {
    if ($$("updateEquipmentForm").validate()) {
        let formValues = $$("updateEquipmentForm").getValues();
        let row = $$("myList").getSelectedItem();
        let equipment = new Equipment();

        row.name = formValues.name;
        row.inventoryNumber = formValues.inventoryNumber;
        let promise = equipment.update(row);

        promise.then(response => {
            return response.json();
        }).then(result => {
                if (result.Error == "") {
                    let datatable = $$("myList");
                    datatable.updateItem(result.Data.id, result.Data);
                    $$(datatable).clearValidation();
                    updateEquipmentForm.hide();
                } else {
                    webix.message(result.Error);
                }
            }
        );
    }
}

//отмена добавления оборудования
function cancel() {
    $$("myList").clearValidation();
    updateEquipmentForm.hide();
}