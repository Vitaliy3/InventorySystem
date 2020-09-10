import {Equipment} from '../../../models/MEquipmentM.js';
import {addEquipmentForm} from './addEquipmentForm.js';
import {updateEquipmentForm} from './updateEquipmentForm.js';
import {RegproductsTree, TreeDatatable} from '../../const.js';
import {hide} from '../recordEquipments.js';

export const equipmentsToolbar = {
    view: "toolbar",
    hidden: hide,
    id: "myToolbar",
    cols: [
        {
            view: "template",
            width: 240,
            css: {"opacity": "0"},
            template: "<div></div>"
        },
        {
            view: "button",
            id: "addEquipment",
            value: "Добавить оборудование",
            height: 50,
            click: addProduct
        },
        {
            view: "button",
            id: "updateEquipment",
            value: "Редактировать оборудование",
            height: 50,
            align: "",
            click: updateProduct
        },
        {
            view: "button",
            id: "deleteEquipment",
            value: "Удалить оборудование",
            height: 50,
            align: "",
            click: deleteProduct
        },
        {
            view: "button",
            id: "writeProduct",
            value: "Списать оборудование",
            height: 50,
            align: "",
            click: writeProduct
        },
    ],
}


//добавление оборудования
function addProduct() {
    let row = $$(RegproductsTree).getSelectedItem();
    if (row) {
        if (row.$level == 3) {
            addEquipmentForm.show();

        } else {
            webix.message("not selected tree item lvl-2");
        }
    } else {
        webix.message("not selected tree item lvl-1");
    }
}

//обновить данные пользователя
function updateProduct() {
    let row = $$(TreeDatatable).getSelectedItem();
    if (row) {
        $$('updateEquipmentForm').setValues({
            name: row.name,
            inventoryNumber: row.inventoryNumber
        });
        updateEquipmentForm.show();
    } else {
        webix.message("not selected item");
    }
}

function deleteProduct() {
    webix.confirm({
        title: "Удаление оборудования",
        text: "Вы уверены?"
    }).then(() => {

        let row = $$(TreeDatatable).getSelectedItem();
        if (row) {
            let product = new Equipment({});
            let promise = product.deleteEquipment(row);
            promise.then(response => {
                return response.json();
            }).then(result => {
                if (result.Error == "") {
                    console.log(result.Data);
                    $$(TreeDatatable).remove(result.Data.id);
                } else {
                    webix.message(result.Error);
                }
            });
        } else {
            webix.message("not selected item");
        }
    })
}

function writeProduct() {
    webix.confirm({
        title: "Списание оборудования",
        text: "Вы уверены?"
    }).then(() => {

        let datatable = $$(TreeDatatable);
        let row = datatable.getSelectedItem();
        let equipment = new Equipment();
        let promise = equipment.writeProduct(row);
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Reject == null) {
                console.log(result.Data);
                datatable.updateItem(result.Data.id, {status: result.Data.status});
            } else {
                console.log("reject null");
            }
        });
    });
}