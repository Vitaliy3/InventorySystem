import {MoveEquipDatatable, MoveEquipmentTree} from '../../const.js';
import {getNeedEquipments} from "../equipment/recordEquipments.js";
import {dragEquipmentTable} from "./components/dragEquipment.js";

//представление: перемещение оборудования
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
                            getNeedEquipments(item, MoveEquipmentTree);
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

