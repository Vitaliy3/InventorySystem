import { Product } from '../../../models/ProductModel.js';
import { RegproductsTree, TreeDatatable } from './../../const.js';

export const addItemForm = webix.ui({
    view: "window",
    width: 600,
    height: 500,
    head: "Добавление элемента",
    autofit: false,
    body: {
        view: "form",
        id: "addItemForm",
        scroll: false,
        width: 600,
        elements: [
            { view: "text", name: "name", label: "Название", labelWidth: 90 },
            { view: "text", name: "inventoryNumber", label: "Инвентарный номер", labelWidth: 150 },
            {
                margin: 5, cols: [
                    { view: "button", label: "Добавить", type: "form", click: addProduct },
                    { view: "button", label: "Отмена", click: closeForm }
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
function addProduct() {
    if ($$("addItemForm").validate()) {
        let formValues = $$("addItemForm").getValues();

        let myTree = $$(RegproductsTree);
        let item = myTree.getSelectedItem();//get class and subclass
        let product = new Product();
        formValues.class = myTree.getItem(item.$parent).class;
        formValues.subclass = item.subclass;
        let promise = product.addProduct(formValues);
        console.log("PROMISE",promise);
        promise.then(
            response => {
                console.log("resp",response);
                $$(TreeDatatable).add(response);
                $$(TreeDatatable).refreshColumns();
                webix.message("success add");
                $$('addItemForm').clear();
                $$('addItemForm').clearValidation();
            },
            err => {
                alert("err" + err);
            });
    }
}

function closeForm() {
    addItemForm.hide();
}
