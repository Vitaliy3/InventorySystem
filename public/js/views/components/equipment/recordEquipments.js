import {RegproductsTree as RegProductsTree, TreeDatatable} from '../../const.js';

//устаналивает класс и подкласс в зависимости от того на каком уровне дерева мы находимся
export function getNeedEquipments(item, id) {
    if (item.$level == 1) {
        document.getElementById("myfilterClass").value = "";
        document.getElementById("myfilterSubclass").value = "";
    }
    if (item.$level == 2) {
        document.getElementById("myfilterClass").value = item.class;
        document.getElementById("myfilterSubclass").value = "";
    }
    if (item.$level == 3) {
        document.getElementById("myfilterClass").value = $$(id).getItem(item.$parent).class;
        document.getElementById("myfilterSubclass").value = item.subclass;
    }
}
//представление: древовидная структура и список оборудования
export const recordEquipments = {
    cols: [
        {
            rows: [
                {
                    view: "tree",
                    id: RegProductsTree,
                    width: 250,
                    columns: [
                        {id: "name", class: "class", fillspace: true,},
                    ],
                    select: "true",
                    on: {
                        onSelectChange: function () {
                            let item = $$(RegProductsTree).getSelectedItem();
                            getNeedEquipments(item, RegProductsTree);
                            $$(TreeDatatable).filterByAll();
                        }
                    }
                },

            ]
        },
        {
            rows: [
                {
                    view: "datatable",
                    id: TreeDatatable,
                    editable: true,
                    scrollY: true,
                    select: true,
                    columns: [
                        {id: "name", class: "class", header: ["Название", {content: "selectFilter"}], fillspace: true,},
                        {
                            id: "user",
                            header: ["Сотрудник", {content: "selectFilter"}],
                            auttowidth: true,
                            fillspace: true
                        },
                        {
                            id: "status",
                            header: ["Статус", {content: "selectFilter"}],
                            auttowidth: true,
                            fillspace: true
                        },
                        {
                            id: "inventoryNumber",
                            header: ["Инвентарный номер", {content: "textFilter"}],
                            fillspace: true
                        },
                    ]
                }]
        }
    ]
};



