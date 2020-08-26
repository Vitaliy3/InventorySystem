
import { toolbar } from './dependViews/toolBar.js';
import { tree } from './dependViews/tree.js';
import { usersList } from './dependViews/UsersView.js';
import { UsersToolbar } from "./dependViews/userToolbar.js";
import { TreeDatatable, UsersDatatable, UserEventsDatatable, MoveProdDatatable, MoveProductTree, DragProdDatatable } from './const.js';
import { Product } from '../models/ProductModel.js';
import { User } from '../models/UserModel.js';
import { userEvents, UserEventToobar } from './userEventsView.js';
import { InventoryEvent } from '../models/UserEvent.js';
import { hide } from './dependViews/tree.js';
import { RegproductsTree } from './const.js';
import { movingTree, moveToolbar } from './dependViews/moveProducts.js';
import { authorizeForm } from './dependViews/forms/authorization.js';

const mainPage = {
    width: 200,
    header: "TESTING",
    height: 1000,
    id: "tabView",
    view: "tabview",
    tabbar: {
        on: {
            onItemClick: function () {
                loadData(this.getValue());
            }
        }
    },

    cells: [
        {
            header: "Учет оборудования",
            id: "regProducts",
            rows: [
                toolbar,
                tree,
            ]
        },
        {
            hidden: hide,
            header: "Учет сотрудников",
            id: "regUsers",
            rows: [
                UsersToolbar,
                usersList,
            ]
        },
        {
            hidden: hide,
            header: "Учет выдачи событий",
            id: "regUserEvents",
            rows: [
                UserEventToobar,
                userEvents
            ]
        },
        {
            hidden: hide,
            header: "Перемещение оборудования",
            id: "moveProducts",
            rows: [
                moveToolbar,
                movingTree,
            ],
        },
    ],
};

//загрузка данных в древовидную таблицу
function pushToTree(id) {
    if (hide) {
        let product = new Product({});
        let promise = product.getUserClasses();
        promise.then(
            result => {
                $$(id).parse(result);
            }
        )
    } else {
        let product = new Product({});
        let promise = product.getAllClasses();
        promise.then(
            result => {
                $$(id).parse(result);
            }
        )
    }
}

//зазрука при переходе на панели
function loadData(id) {
    if (id == "regProducts") {
        pushToTree(RegproductsTree);//parse Tree
        $$(TreeDatatable).clearAll();
        let promise = "";

        //if (hide) - сотрудник
        if (!hide) {
            let product = new Product({});
            promise = product.getAllProducts();
        } else {
            let product = new Product({ id: 1, user: "User" });
            console.log(product);
            promise = product.getUserProducts();
        }
        promise.then(
            result => {
                $$(TreeDatatable).parse(result);
                $$(TreeDatatable).filterByAll();
            },
            err => {
                webix.message("err " + err);
            });
    }

    if (id == "regUsers") {
        $$(UsersDatatable).clearAll();
        let user = new User({});
        let promise = user.getAllUsers(UsersDatatable);
        promise.then(
            result => {
                $$(UsersDatatable).parse(result);
                $$(UsersDatatable).filterByAll();
            },
            err => {
                webix.message("err " + err);
            });
    }

    if (id == "regUserEvents") {
        $$(UserEventsDatatable).clearAll();
        let event = new InventoryEvent({});
        let promise = event.getAllEvents();
        promise.then(
            result => {
                $$(UserEventsDatatable).parse(result);
                $$(UserEventsDatatable).filterByAll();
            },
            err => {
                webix.message("err " + err);
            });
    }

    if (id == "moveProducts") {
        pushToTree(MoveProductTree);//parse Tree
        $$(MoveProductTree).clearAll();
        let promise = "";
        if (!hide) {
            let product = new Product({});
            promise = product.getProdutsInStore();
        }
        promise.then(
            result => {
                $$(MoveProdDatatable).parse(result);
                $$(MoveProdDatatable).filterByAll();
            }
        );
    }
}

webix.ui({
    rows: [
        { view: "button", id: "authorize", value: "Войти", width: 200, height: 50, align: "", click: authorize },//временно расположена здесь 
        mainPage,
    ]
});

//окно авторизации
function authorize() {
    authorizeForm.show({ x: 400, y: 100 });
}
//спиннеры для загрузки
webix.extend($$(TreeDatatable), webix.ProgressBar);
webix.extend($$(UsersDatatable), webix.ProgressBar);
webix.extend($$(UserEventsDatatable), webix.ProgressBar);

//фильтр для для выборки элементов всех подклассов класса в древовидном списке
$$(TreeDatatable).registerFilter(document.getElementById("myfilterClass"),
    { columnId: "class" },
    {
        getValue: function (node) {
            return node.value;
        },
        setValue: function (node, value) {
            node.value = value;
        }
    });

//фильтр для для выборки элементов подкласса в древовидном списке
$$(TreeDatatable).registerFilter(document.getElementById("myfilterSubclass"),
    { columnId: "subclass" },
    {
        getValue: function (node) {
            return node.value;
        },
        setValue: function (node, value) {
            node.value = value;
        }
    });

//фильтр для для выборки элементов всех подклассов класса в древовидном списке
$$(MoveProdDatatable).registerFilter(document.getElementById("myfilterClass"),
    { columnId: "class" },
    {
        getValue: function (node) {
            return node.value;
        },
        setValue: function (node, value) {
            node.value = value;
        }
    });

//фильтр для для выборки элементов подкласса в древовидном списке
$$(MoveProdDatatable).registerFilter(document.getElementById("myfilterSubclass"),
    { columnId: "subclass" },
    {
        getValue: function (node) {
            return node.value;
        },
        setValue: function (node, value) {
            node.value = value;
        }
    });

//событие на drap-grop из склада к сотруднику
$$(DragProdDatatable).attachEvent("onBeforeDrop", function (context, ev) {
    let dnd = webix.DragControl.getContext();
    let value = dnd.from.getItem(dnd.source[0]);
    let product = new Product({ id: value.id, name: value.name, user: value.user });
    console.log(product);
    let result = product.dragToUser();
    return result;
});
//событие на drap-grop от сотрудника на склад
$$(MoveProdDatatable).attachEvent("onBeforeDrop", function (context, ev) {
    let dnd = webix.DragControl.getContext();
    let value = dnd.from.getItem(dnd.source[0]);
    let product = new Product({ id: value.id, name: value.name, user: value.user });
    console.log(product);
    let result = product.dragToStore();
    return result;
});
