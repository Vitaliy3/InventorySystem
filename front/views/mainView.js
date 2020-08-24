
import { toolbar } from './dependViews/toolBar.js';
import { tree } from './dependViews/tree.js';
import { usersList } from './dependViews/UsersView.js';
import { UsersToolbar } from "./dependViews/userToolbar.js";
import { TreeList, UsersList, UserEvent } from './const.js';
import { Product } from '../models/ProductModel.js';
import { User } from '../models/UserModel.js';
//import { authorizeForm } from './dependViews/forms/authorization.js';
import { userEvents } from './userEvents.js';
import { InventoryEvent } from '../models/UserEvent.js';

const mainPage = {
    width: 200,
    header: "TESTING",
    height: 1000,
    id: "tabView",
    view: "tabview",
    tabbar: {
        on: {
            onChange: function () {
                let user = new User({});
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
            header: "Учет сотрудников",
            id: "regUsers",

            rows: [
                UsersToolbar,
                usersList,
            ]
        },
        {
            header: "Учет выдачи событий",
            id: "regUserEvents",
            rows: [
                userEvents
            ]
        },

    ],


};

//анимация загрузки
function loadData(id) {
    if (id == "regProducts") {
        $$(TreeList).clearAll();
        let product = new Product({});
        let promise = product.getAllProducts();
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
        $$(TreeList).clearAll();
        let user = new User({});
        let promise = user.getAllUsers();
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
}

webix.ui({
    rows: [
        mainPage,

    ]
});

//add
webix.extend($$(TreeList), webix.ProgressBar);
webix.extend($$(UsersList), webix.ProgressBar);
webix.extend($$(UserEvent), webix.ProgressBar);


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
