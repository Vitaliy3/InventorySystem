import { Product } from '../../models/ProductModel.js';
import { addItemForm } from './forms/addItemForm.js';
import { updateItemForm } from './forms/updateItemForm.js';
import { Tree, TreeList } from './../const.js';
import { authorizeForm } from './forms/authorization.js';
export const toolbar = {
  view: "toolbar",
  id: "myToolbar",
  cols: [
    { view: "button", id: "addProduct", value: "Добавить оборудование", width: 200, height: 50, click: addProduct },
    { view: "button", id: "updateProduct", value: "Изменить название оборудования", width: 300, height: 50, align: "", click: updateProduct },
    { view: "button", id: "deleteProduct", value: "Удалить оборудование", width: 200, height: 50, align: "", click: deleteProduct },
    { view: "button", id: "writeProduct", value: "Списать оборудование", width: 200, height: 50, align: "", click: writeProduct },
    { view: "button", id: "authorize", value: "Войти", width: 200, height: 50, align: "", click: authorize },//временно расположена здесь 
  ],
}

function authorize() {
  authorizeForm.show({ x: 400, y: 100 });
}
//добавление оборудования
function addProduct() {
  let row = $$(Tree).getSelectedItem();
  if (row) {
    if (row.$level == 3) {
      addItemForm.show({ x: 400, y: 200 });
    } else {
      webix.message("not selected tree item lvl-2");
    }
  } else {
    webix.message("not selected tree item lvl-1");
  }
}

//обновить данные пользователя
function updateProduct() {
  let row = $$(TreeList).getSelectedItem();
  if (row) {
    $$('updateItemForm').setValues({
      name: row.name
    });
    updateItemForm.show({ x: 400, y: 200 });
  } else {
    webix.message("not selected item");
  }
}
function deleteProduct() {
  let row = $$(TreeList).getSelectedItem();
  if (row) {
    let product = new Product(row);
    let promise = product.deleteProduct();
    promise.then(
      result => {
        $$(TreeList).remove(result.id);
      },
      err => {
        alert("err" + err);
      });
  } else {
    webix.message("not selected item");
  }
}
function writeProduct() {
  let row = $$(TreeList).getSelectedItem();
  let product = new Product(row);
  let promise = product.writeProduct();
  promise.then(
    result => {
      let datatable = $$(TreeList);
      datatable.updateItem(result.id, result)
    },
    err => {
      alert("err" + err);
    });
}