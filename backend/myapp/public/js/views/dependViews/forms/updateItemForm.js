import {Equipment} from '../../../models/EquipmentModel.js';

export const updateItemForm = webix.ui({
    view: "window",
    width: 600,
    height: 500,
    head: "Изменение элемента",
    autofit: false,
    body: {
        view: "form",
        id: "updateItemForm",
        scroll: false,
        width: 600,
        elements: [
            {id: "name", view: "text", name: "name", label: "Название"},
            {id: "number", view: "text", name: "inventoryNumber", label: "Инвентарный номер"},

            {
                margin: 5, cols: [
                    {view: "button", label: "Подтвердить", type: "form", click: updateItem},
                    {view: "button", label: "Отмена", click: closeForm}
                ]
            }]
    }
});

function updateItem() {
    let formValues = $$("updateItemForm").getValues();
    let row = $$("myList").getSelectedItem();
    row.name = formValues.name;
    row.inventoryNumber = formValues.inventoryNumber;
    let product = new Equipment();
    let promise = product.updateEquipment(row);
    promise.then(response => {
        return response.json();
    }).then(result => {
            if (result.Reject == null) {
                let datatable = $$("myList");
                datatable.updateItem(row.id, result.Data);
                updateItemForm.hide();
            } else {
                webix.message(result.Reject);
            }
        }
    );
}

function closeForm() {
    updateItemForm.hide();
}