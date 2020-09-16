import {Equipment} from '../../../models/MEquipmentM.js';

export const updateEquipmentForm = webix.ui({
    view: "window",
    width: 600,
    position:"center",
    height: 500,
    head: "Редактирование оборудования",
    autofit: false,
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
            }]
    }
});

function confirmUpdateItem() {
    let formValues = $$("updateEquipmentForm").getValues();
    let row = $$("myList").getSelectedItem();
    row.name = formValues.name;
    row.inventoryNumber = formValues.inventoryNumber;
    let product = new Equipment();
    let promise = product.updateEquipment(row);
    promise.then(response => {
        return response.json();
    }).then(result => {
            if (result.Error == "") {
                let datatable = $$("myList");
                datatable.updateItem(result.Data.id, result.Data);
                updateEquipmentForm.hide();
            } else {
                webix.message(result.Error);
            }
        }
    );
}

function cancel() {
    updateEquipmentForm.hide();
}