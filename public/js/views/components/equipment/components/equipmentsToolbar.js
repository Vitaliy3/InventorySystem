import {Equipment} from '../../../../models/MEquipmentM.js';
import {addEquipmentForm} from './addEquipmentForm.js';
import {updateEquipmentForm} from './updateEquipmentForm.js';
import {RegproductsTree, TreeDatatable} from '../../../const.js';

//панель инструментов для учета оборудования
export const equipmentsToolbar = {
    view: "toolbar",
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
            click: addEquipment
        },
        {
            view: "button",
            id: "updateEquipment",
            value: "Редактировать оборудование",
            height: 50,
            align: "",
            click: updateEquipment
        },
        {
            view: "button",
            id: "deleteEquipment",
            value: "Удалить оборудование",
            height: 50,
            align: "",
            click: deleteEquipment
        },
        {
            view: "button",
            id: "writeProduct",
            value: "Списать оборудование",
            height: 50,
            align: "",
            click: writeEquipment
        },
    ],
};

//добавление оборудования
function addEquipment() {
    let row = $$(RegproductsTree).getSelectedItem();

    if (row) {
        if (row.$level == 3) {
            addEquipmentForm.show();

        } else {
            webix.message("Не выбран подкласс");
        }
    } else {
        webix.message("Не выбран класс");
    }
}

//редактирование пользователя
function updateEquipment() {
    let row = $$(TreeDatatable).getSelectedItem();

    if (row) {
        $$('updateEquipmentForm').setValues({
            name: row.name,
            inventoryNumber: row.inventoryNumber
        });
        updateEquipmentForm.show();
    } else {
        webix.message("Оборудование не выбрано");
    }
}

//удаление оборудования
function deleteEquipment() {
    let row = $$(TreeDatatable).getSelectedItem();
    if (row) {
        webix.confirm({
            title: "Удаление оборудования",
            text: "Вы уверены?"
        }).then(() => {
            let product = new Equipment();
            let promise = product.delete(row);
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
        });
    } else {
        webix.message("Оборудование не выбрано");
    }
}

//списание оборудования
function writeEquipment() {
    let datatable = $$(TreeDatatable);
    let row = datatable.getSelectedItem();
    if (row) {
        webix.confirm({
            title: "Списание оборудования",
            text: "Вы уверены?"
        }).then(() => {
            let equipment = new Equipment();
            let promise = equipment.write(row);
            promise.then(response => {
                return response.json();
            }).then(result => {
                if (result.Error == "") {
                    datatable.updateItem(result.Data.id, {status: result.Data.status});
                } else {
                    webix.message(result.Error);
                }
            });
        });
    } else {
        webix.message("Оборудование не выбрано");
    }
}
