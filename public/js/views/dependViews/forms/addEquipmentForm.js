import {Equipment} from '../../../models/MEquipmentM.js';
import {RegproductsTree, TreeDatatable} from './../../const.js';

export const addEquipmentForm = webix.ui({
    view: "window",
    width: 600,
    position:"center",
    height: 500,
    head: "Добавление оборудования",
    autofit: false,
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
                    {view: "button", label: "Отмена", click: closeForm}
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

function confirmAddEuipment() {
    if ($$("addEquipmentForm").validate()) {
        let formValues = $$("addEquipmentForm").getValues();

        let myTree = $$(RegproductsTree);
        let item = myTree.getSelectedItem();//get class and subclass
        let product = new Equipment();
        formValues.class = myTree.getItem(item.$parent).class;
        formValues.subclass = item.subclass;
        let promise = product.addEquipment(formValues);
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                console.log("data",result.Data);
                $$(TreeDatatable).add(result.Data);
                $$(TreeDatatable).refreshColumns();
                webix.message("success add");
                $$('addEquipmentForm').clear();
                $$('addEquipmentForm').clearValidation();
            } else {
                webix.message(result.Error)
            }
        })
    }
}

function closeForm() {
    addEquipmentForm.hide();
}
