
import { MoveTree, MoveProdTree, MoveProduct } from './../const.js';
import { getNeedProducts } from './tree.js';
import { User } from '../../models/UserModel.js';
import { Product } from '../../models/ProductModel.js';
export var new_options = [];

let user = new User({});
let users = user.getAllUsers("");
users.then(result => {
    let joinUsers = [];
    let temp = "";
    for (let i = 0; i < result.length; i++) {
        temp = { name: result[i].id + ". " + result[i].name + " " + result[i].surname + " " + result[i].patronymic };
        joinUsers.push(temp);
    }
    console.log(joinUsers);
    let list = $$("combo").getPopup().getList();
    list.clearAll();
    list.parse(joinUsers);
});

export const moveToolbar = {
    view: "toolbar",
    id: "moveToolbar",
    cols: [
        {
            view: "combo",
            value: 2,
            id: "combo",
            options: {
                body: {
                    template: "#name#"
                },
                data: []
            }
        },
        // { view: "select", id: "selectU", value: 2, label: "Поиск товара сотрудника", labelwidth: 200, options: userData },
        { view: "button", value: "Найти", click: filterUsers, width: 100 },
    ],
}

export const dragProductTable = {
    header: "TEST",
    view: "datatable",
    drag: true,
    id: MoveProduct,
    width: 500,
    select: true,
    columns: [
        { id: "name", header: "Название", class: "class", fillspace: true, },
        { id: "inventoryNumber", header: "Инвентарный номер", fillspace: true },
    ]
};

function filterUsers() {
    let selected = $$("combo").getText();
    let split = selected.split(".");
    console.log(split[0]);
    let product = new Product({ id: split[0], user: "User1" });
    let promise = product.getProductsUser();
    promise.then(result => {
        $$(MoveProduct).parse(result);
    })
}

export const movingTree = {
    header: "TEST",
    cols: [
        {
            rows: [
                {
                    view: "tree",
                    id: MoveTree,
                    width: 250,
                    columns: [
                        { id: "name", class: "class", fillspace: true, },
                    ],
                    select: "true",
                    on: {
                        onSelectChange: function () {
                            let item = $$(MoveTree).getSelectedItem();
                            $$(MoveProdTree).parse(getNeedProducts(item, MoveTree));
                            $$(MoveProdTree).filterByAll();//refresh data after change tree column
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
                    id: MoveProdTree,
                    editable: true,
                    drag: true,
                    editaction: "custom",
                    select: true,
                    columns: [
                        { id: "name", header: "Название", class: "class", fillspace: true, },
                        { id: "status", header: "Статус", auttowidth: true, fillspace: true },
                        { id: "inventoryNumber", header: "Инвентарный номер", fillspace: true },
                    ]
                },

            ],

        },

        dragProductTable,

    ]
};


