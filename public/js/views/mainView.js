import {equipmentsToolbar} from './dependViews/forms/equipmentsToolbar.js';
import {recordEquipments} from './dependViews/recordEquipments.js';
import {usersList} from './dependViews/recordUsers.js';
import {UsersToolbar} from "./dependViews/forms/employeeToolbar.js";
import {
    combo,
    DragProdDatatable,
    MoveEquipDatatable,
    MoveEquipmentTree,
    RegproductsTree,
    TreeDatatable,
    UserEventsDatatable,
    UsersDatatable
} from './const.js';
import {Equipment} from '../models/MEquipmentM.js';
import {Employee} from '../models/MUserModel.js';
import {recordEvents, recordEventsDatapicker} from './dependViews/recordEvents.js';
import {UserEvent} from '../models/MEmployeeEvents.js';
import {moveToolbar, movingTree} from './dependViews/equipmentsMove.js';

const mainPage = {
    width: 200,
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
                equipmentsToolbar,
                recordEquipments,
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
                recordEventsDatapicker,
                recordEvents
            ]
        },
        {
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
    let role = getUserRole();
    let promise = "";
    let product = new Equipment();
    if (role != "admin") {
        let token = getCurrentToken();
        promise = product.getFullTree(token);
    } else {
        promise = product.getFullTree("");
    }
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            $$(id).parse(result.Data);
        } else {
            webix.message(result.Error);
        }
    });
}

//зазрука при переходе на панели
function loadData(id) {
    if (id == "regProducts") {
        pushToTree(RegproductsTree);//parse Tree
        $$(TreeDatatable).clearAll();
        let promise = "";

        let product = new Equipment();
        let role = getUserRole();
        if (role == "admin") {
            promise = product.getAllEquipment("");
        } else {
            let token = getCurrentToken();

            promise = product.getAllEquipment(token);
        }
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(TreeDatatable).parse(result.Data);
                $$(TreeDatatable).filterByAll();
            } else {
                webix.message("err", result.Error);
            }
        });
    }

    if (id == "regUsers") {
        $$(UsersDatatable).clearAll();
        let user = new Employee();
        let promise = user.getAllEmployees();
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(UsersDatatable).parse(result.Data);
                $$(UsersDatatable).filterByAll();
            } else {
                webix.message(result.Error);
            }
        })
    }

    if (id == "regUserEvents") {
        $$(UserEventsDatatable).clearAll();
        let event = new UserEvent();
        let promise = event.getAllEvents();
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(UserEventsDatatable).clearAll();
                $$(UserEventsDatatable).parse(result.Data);

                $$(UserEventsDatatable).filterByAll();
            } else {
                webix.message(result.Error);
            }
        })
    }

    if (id == "moveProducts") {
        pushToTree(MoveEquipmentTree);//parse Tree
        $$(MoveEquipmentTree).clearAll();
        let promise = "";
        let product = new Equipment({});
        promise = product.getEquipmentsInStore();
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                console.log(result.Data);
                $$(MoveEquipDatatable).parse(result.Data);
                $$(MoveEquipDatatable).filterByAll();
            } else {
                webix.message(result.Error);
            }
        });
        //заполнение select
        let user = new Employee();
        let users = user.getAllEmployees();
        users.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                let joinUsers = [];
                let temp = "";
                for (let i = 0; i < result.Data.length; i++) {
                    temp = {
                        id: result.Data[i].id,
                        name: result.Data[i].name + " " + result.Data[i].surname + " " + result.Data[i].patronymic
                    };
                    joinUsers.push(temp);
                }
                let list = $$(combo).getPopup().getList();
                list.clearAll();
                list.parse(joinUsers);
                console.log(joinUsers);
            } else {
                webix.message(result.Error);
            }
        });
    }
}

webix.ui({
    rows: [
        {view: "button", id: "authorize", value: "Выйти", width: 200, height: 50, align: "right", click: logout},//временно расположена здесь
        mainPage,
    ]
});

function getCookie(name) {
    let matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}

function getUserRole() {
    let cookie = getCookie("token");
    let splitCookie = cookie.split(':');
    return splitCookie[1];
}

function getCurrentToken() {
    let cookie = getCookie("token");
    let splitCookie = cookie.split(':');
    return splitCookie[0];
}


window.onload = function () {
    let cookie = getCookie("token");
    let splitCookie = cookie.split(':');
    if (splitCookie[1] != "admin") {
        $$("tabView").getTabbar().removeOption("moveProducts");
        $$("tabView").getTabbar().removeOption("regUsers");
        $$("tabView").getTabbar().removeOption("regUserEvents");
        $$("myToolbar").hide();
        $$("myList").hideColumn("status");
        $$("myList").hideColumn("user");
    }
};

function logout() {
    let promise = fetch("/logout")
    promise.then(json => {
        return json.json()
    }).then(result => {
        if (result.Error == "") {
            document.cookie = "auth" + '=; Max-Age=0';
            document.cookie = "token" + '=; Max-Age=0';
            document.location.href = result.Data;
        } else {
            webix.message(result.Error);
        }
    });
    //  authorizeForm.show({x: 400, y: 100});  --DEL
}

//спиннеры для загрузки
webix.extend($$(TreeDatatable), webix.ProgressBar);
webix.extend($$(UsersDatatable), webix.ProgressBar);
webix.extend($$(UserEventsDatatable), webix.ProgressBar);
$$("write")

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
$$(MoveEquipDatatable).registerFilter(document.getElementById("myfilterClass"),
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
$$(MoveEquipDatatable).registerFilter(document.getElementById("myfilterSubclass"),
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
    let product = new Equipment(value);
    let selected = $$(combo).getValue();
    if (selected == 2) {
        webix.message("Не выбран пользователь");
        return false;
    }
    let sendValue = {fk_user: selected, id: value.id};
    sendValue.fk_user = parseInt(sendValue.fk_user);
    let promise = product.dragToUser(sendValue);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            return result.Data;
        } else {
            webix.message(result.Error);
            webix.message("Reload page");
            return false;
        }
    });
    return true;
});

//событие на drap-grop от сотрудника на склад
$$(MoveEquipDatatable).attachEvent("onBeforeDrop", function (context, ev) {
    let dnd = webix.DragControl.getContext();
    let value = dnd.from.getItem(dnd.source[0]);
    let product = new Equipment();
    let promise = product.dragToStore(value);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
        } else {
            webix.message(result.Error);
        }
    });
    return true;
});
$$(MoveEquipDatatable).attachEvent("onAfterAdd", function (id, index) {
    $$(MoveEquipDatatable).filterByAll();
});


