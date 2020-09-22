import {Equipment} from '../../../../models/MEquipmentM.js';
import {RegproductsTree, TreeDatatable} from '../../../const.js';

//форма для добавления оборудования
export const addEquipmentForm = webix.ui({
    view: "window",
    width: 600,
    position:"center",
    height: 500,
    head: "Добавление оборудования",
    body: {
        view: "form",
        id: "addEquipmentForm",
        scroll: false,
        width: 600,
        elements: [
            {view: "text", name: "name", label: "Название", labelWidth: 90},
            {view: "text", name: "inventoryNumber", label: "Инвентарный номер", labelWidth: 150},
            {
                margin: 5, cols: [
                    {view: "button", label: "Добавить", type: "form", click: confirmAddEuipment},
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
        }
    }
});

///вызывается при подтверждении добавления оборудования,считывает данные из формы,отправляет запрос на добавление
function confirmAddEuipment() {
    if ($$("addEquipmentForm").validate()) {
        let formValues = $$("addEquipmentForm").getValues();
        let myTree = $$(RegproductsTree);
        let item = myTree.getSelectedItem();//get class and subclass
        let product = new Equipment();
        let promise = product.add(formValues);

        formValues.class = myTree.getItem(item.$parent).class;
        formValues.subclass = item.subclass;
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                console.log("data",result.Data);
                $$(TreeDatatable).add(result.Data);
                webix.message("Обрудование добавлено");

                $$('addEquipmentForm').clear();
                $$('addEquipmentForm').clearValidation();
            } else {
                webix.message(result.Error)
            }
        })
    }
}

//отмена в форме добавления оборудования
function cancel() {
    $$('addEquipmentForm').clearValidation();
    addEquipmentForm.hide();
}
