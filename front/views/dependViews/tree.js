import { Product } from '../../models/ProductModel.js';
let product = new Product({});
let treeData = product.getClassSubclass();
let productData = product.getAllProducts();
function getNeedProducts(item) { //возр оборудование в классе/подклассе
    let result = [];
    if (item.$level == 1) {
        for (let i = 0; i < productData.length; i++) {
            if (item.class == productData[i].fkClass)
                result.push(productData[i]);
        }
    }
    else {
        let parent = $$("myTree").getItem(item.$parent);
        for (let i = 0; i < productData.length; i++) {
            if (parent.class == productData[i].fkClass && item.subclass == productData[i].fk_subClass)
                result.push(productData[i]);
        }
    }
    console.log(result);
    return result;
}
export const tree = {
    cols: [

        {
            rows: [
                {
                    view: "tree",
                    id: "myTree",
                    width: 250,
                    data: treeData,
                    select: "true",
                    on: {
                        onSelectChange: function () {
                            let item = $$("myTree").getSelectedItem();
                            $$("myList").clearAll();
                            $$("myList").parse(getNeedProducts(item));
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
                    id: "myList",
                    editable: true,
                    editaction: "custom",
                    select: true,
                    data: productData,
                    columns: [
                        { id: "name", header: ["Название", { content: "selectFilter" }], width: 200, },
                        { id: "user", header: ["Сотрудник", { content: "selectFilter" }], width: 200 },
                        { id: "status", header: ["Статус", { content: "selectFilter" }], width: 100 },
                        { id: "inventoryNumber", header: ["Инвентарный номер", { content: "selectFilter" }], width: 200 },
                    ]
                }]
        }
    ]
};



