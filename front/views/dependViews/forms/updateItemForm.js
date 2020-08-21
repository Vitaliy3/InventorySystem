import { Product } from '../../../models/ProductModel.js';
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
            { id: "b1", view: "text", name: "name", label: "Название" },
            {
                margin: 5, cols: [
                    { view: "button", label: "Подтвердить", type: "form", click: updateItem },
                    { view: "button", label: "Отмена", click: closeForm }
                ]
            }]
    }
});
function updateItem() {
    let formValues = $$("updateItemForm").getValues();
    let row = $$("myList").getSelectedItem();
    row.name = formValues.name;//changed value
    let product = new Product(row);
    let promise = product.updateProduct();
    promise.then(
        result => {
            let datatable = $$("myList");
            datatable.updateItem(result.id, result)
            updateItemForm.hide();
        },
        err => {
            alert("err" + err);
        });
}
function closeForm() {
    updateItemForm.hide();
}