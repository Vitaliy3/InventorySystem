import {combo, DragProdDatatable, MoveEquipDatatable, MoveEquipmentTree} from './../const.js';
import {getNeedProducts} from "./recordEquipments.js";
import {Equipment} from "../../models/MEquipmentM.js";

export const moveToolbar = {
    view: "toolbar",
    id: "moveToolbar",
    cols: [
        {
            view: combo,
            value: 2,
            id: combo,
            width: 200,
            align: "right",
            options: {
                body: {

                    template: "<span style='display:none;>#id#</span> <span>#name#</span>"
                },
            }
        },
        {view: "button",value: "Найти", click: filterUsers, width: 100,},
    ],
};

export const dragEquipmentTable = {
    view: "datatable",
    drag: true,
    id: DragProdDatatable,
    width: 500,
    select: true,
    columns: [
        {id: "name", header: "Название", class: "class", fillspace: true,},
        {id: "inventoryNumber", header: "Инвентарный номер", fillspace: true},
    ]
};

function filterUsers() {
    let eqipment = new Equipment();
    let selected = $$(combo).getValue();
    $$(DragProdDatatable).clearAll();
    let promise = eqipment.getUserEquipments(selected);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            $$(DragProdDatatable).parse(result.Data);
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
                            $$(MoveEquipDatatable).filterByAll();//refresh data after change tree column
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

