import { Product } from '../../models/ProductModel.js';
import { Tree, TreeList } from './../const.js';
let product = new Product({});
let treeData = product.getClassSubclass();
function getNeedProducts(item) { //возр оборудование в классе/подклассе
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
        document.getElementById("myfilterClass").value = $$("myTree").getItem(item.$parent).class;
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
                    data: treeData,
                    select: "true",
                    on: {
                        onSelectChange: function () {
                            let item = $$(Tree).getSelectedItem();
                            $$(TreeList).parse(getNeedProducts(item));
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
                        { id: "user", header: ["Сотрудник", { content: "selectFilter" }], auttowidth: true, fillspace: true },
                        { id: "status", header: ["Статус", { content: "selectFilter" }], auttowidth: true, fillspace: true },
                        { id: "inventoryNumber", header: ["Инвентарный номер", { content: "textFilter" }], fillspace: true },
                    ]
                }]
        }
    ]
};



