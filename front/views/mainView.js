
import { toolbar } from './dependViews/toolBar.js';
import { tree } from './dependViews/tree.js';
import { usersList } from './dependViews/UsersView.js';
import { UsersToolbar } from "./dependViews/userToolbar.js";
import { TreeList, UsersList, UserEvent, MoveProdTree, MoveTree, MoveProduct } from './const.js';
import { Product } from '../models/ProductModel.js';
import { User } from '../models/UserModel.js';
import { userEvents, UserEventToobar } from './userEventsView.js';
import { InventoryEvent } from '../models/UserEvent.js';
import { hide } from './dependViews/tree.js';
import { Tree } from './const.js';
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
        pushToTree(Tree);//parse Tree
        $$(TreeList).clearAll();
        let promise = "";

        //if (hide) - сотрудник
        if (!hide) {
            let product = new Product({});
            promise = product.getAllProducts();
        } else {
            let product = new Product({ id: 1, user: "User" });
            console.log(product);
            promise = product.getProductsUser();
        }
        promise.then(
            result => {
                $$(TreeList).parse(result);
                $$(TreeList).filterByAll();
            },
            err => {
                webix.message("err " + err);
            });
    }

    if (id == "regUsers") {
        $$(UsersList).clearAll();
        let user = new User({});
        let promise = user.getAllUsers(UsersList);
        promise.then(
            result => {
                $$(UsersList).parse(result);
                $$(UsersList).filterByAll();
            },
            err => {
                webix.message("err " + err);
            });
    }

    if (id == "regUserEvents") {
        $$(UserEvent).clearAll();
        let event = new InventoryEvent({});
        let promise = event.getAllEvents();
        promise.then(
            result => {
                $$(UserEvent).parse(result);
                $$(UserEvent).filterByAll();
            },
            err => {
                webix.message("err " + err);
            });
    }

    if (id == "moveProducts") {
        pushToTree(MoveTree);//parse Tree
        $$(MoveTree).clearAll();
        let promise = "";
        if (!hide) {
            let product = new Product({});
            promise = product.getProdutsInStore();
        }
        promise.then(
            result => {
                $$(MoveProdTree).parse(result);
                $$(MoveProdTree).filterByAll();
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
webix.extend($$(TreeList), webix.ProgressBar);
webix.extend($$(UsersList), webix.ProgressBar);
webix.extend($$(UserEvent), webix.ProgressBar);

//фильтр для для выборки подклассов класса в древовидном списке
$$(TreeList).registerFilter(document.getElementById("myfilterClass"),
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
$$(TreeList).registerFilter(document.getElementById("myfilterSubclass"),
    { columnId: "subclass" },
    {
        getValue: function (node) {
            return node.value;
        },
        setValue: function (node, value) {
            node.value = value;
        }
    });

//фильтр для для выборки подклассов класса в древовидном списке
$$(MoveProdTree).registerFilter(document.getElementById("myfilterClass"),
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
$$(MoveProdTree).registerFilter(document.getElementById("myfilterSubclass"),
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
$$(MoveProduct).attachEvent("onBeforeDrop", function (context, ev) {
    let dnd = webix.DragControl.getContext();
    let value = dnd.from.getItem(dnd.source[0]);
    let product = new Product({ id: value.id, name: value.name, user: value.user });
    console.log(product);
    let result = product.dragToUser();
    return result;
});
//событие на drap-grop от сотрудника на склад
$$(MoveProdTree).attachEvent("onBeforeDrop", function (context, ev) {
    let dnd = webix.DragControl.getContext();
    let value = dnd.from.getItem(dnd.source[0]);
    let product = new Product({ id: value.id, name: value.name, user: value.user });
    console.log(product);
    let result = product.dragToStore();
    return result;
});
