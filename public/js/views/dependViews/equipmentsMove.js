import {combo, DragProdDatatable, MoveEquipDatatable, MoveEquipmentTree} from './../const.js';
import {getNeedProducts} from "./recordEquipments.js";
import {Equipment} from "../../models/MEquipmentM.js";

export const moveToolbar = {
    view: "toolbar",
    id: "moveToolbar",
    width: 200,
    cols: [
        {
            view: "template",
            css: {"opacity": "0"},
            template: "<div></div>"
        },
        {
            view: combo,
            value: 2,
            id: combo,
            width: 295,
            align: "right",
            options: {
                body: {
                    template: "<span style='display:none'>#id#</span> <span>#name#</span>"
                },
            }
        },
        {view: "button", id: "button1", value: "Найти", click: findEquipmentsByUser, width: 100,},
    ],
};

export const dragEquipmentTable = {
    view: "datatable",
    drag: true,
    id: DragProdDatatable,
    width: 400,
    select: true,
    columns: [
        {id: "Class", header: "Класс", fillspace: true,},
        {id: "Subclass", header: "Подкласс", width:150,},
        {id: "name", header: "Название", fillspace: true,},
        {id: "inventoryNumber", header: "Инвентарный номер", fillspace: true},
    ]
};

function findEquipmentsByUser() {
    let eqipment = new Equipment();
    let selected = $$(combo).getValue();
    $$(DragProdDatatable).clearAll();
    console.log(selected);
    let promise = eqipment.getEquipmentsByUser(selected);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            if (result.Data == null) {
            } else {
                console.log("data", result.Data);
                $$(DragProdDatatable).parse(result.Data);
            }
        } else {
            webix.message(result.Error);
        }
    });
}

export const movingTree = {
    cols: [
        {
            rows: [
                {
                    view: "tree",
                    id: MoveEquipmentTree,
                    width: 250,
                    columns: [
                        {id: "name", class: "class", fillspace: true,},
                    ],
                    select: "true",
                    on: {
                        onSelectChange: function () {
                            let item = $$(MoveEquipmentTree).getSelectedItem();
                            $$(MoveEquipDatatable).parse(getNeedProducts(item, MoveEquipmentTree));
                            $$(MoveEquipDatatable).filterByAll();
                        }
                    }
                },
            ]
        },
        {
            rows: [
                {
                    view: "datatable",
                    id: MoveEquipDatatable,
                    editable: true,
                    drag: true,
                    editaction: "custom",
                    select: true,
                    columns: [
                        {id: "name", header: "Название", class: "class", fillspace: true,},
                        {id: "status", header: "Статус", auttowidth: true, fillspace: true},
                        {id: "inventoryNumber", header: "Инвентарный номер", fillspace: true},
                    ]
                },

            ],

        },

        dragEquipmentTable,

    ]
};

