import { Product } from '../../models/ProductModel.js';
import { addItemForm } from './forms/addItemForm.js';
import { updateItemForm } from './forms/updateItemForm.js';
export const toolbar = {
  view: "toolbar",
  id: "myToolbar",
  cols: [
    { view: "button", id: "addProduct", value: "Добавить оборудование", width: 200, height: 50, click: addProduct },
    { view: "button", id: "updateProduct", value: "Изменить название оборудования", width: 300, height: 50, align: "", click: updateProduct },
    { view: "button", id: "deleteProduct", value: "Удалить оборудование", width: 200, height: 50, align: "", click: deleteProduct },
    { view: "button", id: "writeProduct", value: "Списать оборудование", width: 200, height: 50, align: "", click: writeProduct }
  ],
}
//добавление оборудования
function addProduct() {
  let row = $$("myTree").getSelectedItem();
  if (row) {
    if (row.$level == 2) {
      addItemForm.show({ x: 200, y: 200 });
    } else {
      webix.message("not selected tree item lvl-2");
    }
  } else {
    webix.message("not selected tree item lvl-1");
  }
}

//обновить данные пользователя
function updateProduct() {
  let row = $$("myList").getSelectedItem();
  if (row) {
    updateItemForm.show({ x: 200, y: 200 });
    $$('updateItemForm').setValues({
      name: row.name
    });
  } else {
    webix.message("not selected item");
  }
}
function deleteProduct() {
  let row = $$("myList").getSelectedItem();
  if (row) {
    let product = new Product(row);
    let promise = product.deleteProduct();
    promise.then(
      result => {
        $$("myList").remove(result.id);
      },
      err => {
        alert("err" + err);
      });
  } else {
    webix.message("not selected item");
  }
}
function writeProduct() {
  let row = $$("myList").getSelectedItem();
  let product = new Product(row);
  let promise = product.writeProduct();
  promise.then(
    result => {
      let datatable = $$("myList");
      datatable.updateItem(result.id, result)
    },
    err => {
      alert("err" + err);
    });
}