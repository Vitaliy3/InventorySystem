import {toolbar} from './dependViews/toolBar.js';
import {hide, tree} from './dependViews/tree.js';
import {usersList} from './dependViews/UsersView.js';
import {UsersToolbar} from "./dependViews/userToolbar.js";
import {
    DragProdDatatable,
    MoveProdDatatable,
    MoveProductTree,
    RegproductsTree,
    TreeDatatable,
    UserEventsDatatable,
    UsersDatatable
} from './const.js';
import {Equipment} from '../models/EquipmentModel.js';
import {Employee} from '../models/UserModel.js';
import {userEvents, UserEventToobar} from './userEventsView.js';
import {InventoryEvent} from '../models/UserEvent.js';
import {moveToolbar, movingTree} from './dependViews/moveProducts.js';
import {authorizeForm} from './dependViews/forms/authorization.js';

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
    $$(id).clearAll();
    if (hide) {
        let product = new Equipment();
        let promise = product.getUserClasses();
        promise.then(
            result => {
                $$(id).parse(result);
            }
        )
    } else {
        let product = new Equipment({});
        let promise = product.getAllTree();
        promise.then(response => {
            return response.json();
        }).then(result => {
                if (result.Reject == null) {
                    console.log("tree", result.Tree);
                    $$(id).parse(result.Tree);
                } else {
                    webix.message(result.Reject);
                }
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
            let product = new Equipment({});
            promise = product.getAllEquipment();
        } else {
            let product = new Equipment({id: 1, user: "Employee"});
            promise = product.getUserProducts();
        }

        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Reject == null) {
                $$(TreeDatatable).parse(result.DataArray);
                $$(TreeDatatable).filterByAll();
            } else {
                webix.message("err", result.Reject);
            }
        });
    }

    if (id == "regUsers") {
        $$(UsersDatatable).clearAll();
        let user = new Employee({});
        let promise = user.getAllEmployees();
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Reject == null) {
                $$(UsersDatatable).parse(result.DataArray);
                $$(UsersDatatable).filterByAll();
            } else {
                webix.message(result.Reject);
            }
        })
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
            let product = new Equipment({});
            promise = product.getEquipmentsInStore();
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
        {view: "button", id: "authorize", value: "Войти", width: 200, height: 50, align: "", click: authorize},//временно расположена здесь
        mainPage,
    ]
});

//окно авторизации
function authorize() {
    authorizeForm.show({x: 400, y: 100});
}

//спиннеры для загрузки
webix.extend($$(TreeDatatable), webix.ProgressBar);
webix.extend($$(UsersDatatable), webix.ProgressBar);
webix.extend($$(UserEventsDatatable), webix.ProgressBar);

//фильтр для для выборки элементов всех подклассов класса в древовидном списке
$$(TreeDatatable).registerFilter(document.getElementById("myfilterClass"),
    {columnId: "class"},
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
    {columnId: "subclass"},
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
    {columnId: "class"},
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
    {columnId: "subclass"},
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
    let product = new Equipment({id: value.id, name: value.name, user: value.user});
    console.log(product);
    let result = product.dragToUser();
    return result;
});
//событие на drap-grop от сотрудника на склад
$$(MoveProdDatatable).attachEvent("onBeforeDrop", function (context, ev) {
    let dnd = webix.DragControl.getContext();
    let value = dnd.from.getItem(dnd.source[0]);
    let product = new Equipment({id: value.id, name: value.name, user: value.user});
    console.log(product);
    let result = product.dragToStore();
    return result;
});
