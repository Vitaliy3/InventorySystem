
import { Tree, TreeList } from './../const.js';
export const hide = false;

export function getNeedProducts(item, id) { //возр оборудование в классе/подклассе
    let result = [];
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
    return result;
}
export const tree = {
    header: "TEST",
    cols: [
        {
            rows: [
                {
                    view: "tree",
                    id: Tree,
                    width: 250,
                    columns: [
                        { id: "name", class: "class", fillspace: true, },
                    ],
                    select: "true",
                    on: {
                        onSelectChange: function () {
                            let item = $$(Tree).getSelectedItem();
                            $$(TreeList).parse(getNeedProducts(item, Tree));
                            $$(TreeList).filterByAll();//refresh data after change tree column
                        }
                    }
                },

            ]
        },
        { view: "resizer" },
        {
            rows: [
                {
                    view: "datatable",
                    id: TreeList,
                    editable: true,
                    editaction: "custom",
                    select: true,

                    columns: [
                        { id: "name", class: "class", header: ["Название", { content: "selectFilter" }], fillspace: true, },
                        { id: "user", hidden: hide, header: ["Сотрудник", { content: "selectFilter" }], auttowidth: true, fillspace: true },
                        { id: "status", hidden: hide, header: ["Статус", { content: "selectFilter" }], auttowidth: true, fillspace: true },
                        { id: "inventoryNumber", header: ["Инвентарный номер", { content: "textFilter" }], fillspace: true },
                    ]
                }]
        }
    ]
};



