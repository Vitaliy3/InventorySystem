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
        {view: "button", id: "addEquipment", value: "Добавить оборудование", width: 200, height: 50, click: addProduct},
        {
            view: "button",
            id: "updateEquipment",
            value: "Изменить название оборудования",
            width: 300,
            height: 50,
            align: "",
            click: updateProduct
        },
        {
            view: "button",
            id: "deleteEquipment",
            value: "Удалить оборудование",
            width: 200,
            height: 50,
            align: "",
            click: deleteProduct
        },
        {
            view: "button",
            id: "writeProduct",
            value: "Списать оборудование",
            width: 200,
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
            addEquipmentForm.show({x: 400, y: 200});
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
        updateEquipmentForm.show({x: 400, y: 200});
    } else {
        webix.message("not selected item");
    }
}

function deleteProduct() {
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
}

function writeProduct() {
    let datatable = $$(TreeDatatable);
    let row = datatable.getSelectedItem();
    let equipment = new Equipment();
    let promise = equipment.writeProduct(row);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Reject == null) {
            datatable.updateItem(result.Data.id, {status: result.Data.status});
        } else {
            console.log("reject null");
        }
    });
}