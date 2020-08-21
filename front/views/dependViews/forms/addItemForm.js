import { Product } from '../../../models/ProductModel.js';
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
            { view: "text", name: "user", label: "Сотрудник", labelWidth: 90 },
            { view: "text", name: "status", label: "Статус", labelWidth: 90 },
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
function addProduct() {
    if ($$("addItemForm").validate()) {
        let formValues = $$("addItemForm").getValues();

        let myTree = $$("myTree");
        let item = myTree.getSelectedItem();//get class and subclass
        let product = new Product(formValues);
        product.class = myTree.getItem(item.$parent).class;
        product.subclass = item.subclass;

        console.log("PRODUCT:", product);
        let promise = product.addProduct();
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
    }
}

function closeForm() {
    addItemForm.hide();
}
