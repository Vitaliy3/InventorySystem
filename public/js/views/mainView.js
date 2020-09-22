import {equipmentsToolbar} from './components/equipment/components/equipmentsToolbar.js';
import {recordEquipments} from './components/equipment/recordEquipments.js';
import {usersList} from './components/emoloyee/recordEmoloyees.js';
import {UsersToolbar} from "./components/emoloyee/components/toolbar.js";
import {
    combo,
    DragProdDatatable,
    EmployeeEventsDatatable,
    MoveEquipDatatable,
    MoveEquipmentTree,
    RegproductsTree,
    TreeDatatable,
    UsersDatatable
} from './const.js';
import {Equipment} from '../models/MEquipmentM.js';
import {Employee} from '../models/MEmployeeModel.js';
import {recordEventsDatepicker} from './components/event/components/datepickerToolbar.js';
import {recordEvent} from './components/event/recordEvents.js';
import {UserEvent} from '../models/MEmployeeEvents.js';
import {movingTree} from './components/equipmentsMove/equipmentsMove.js';
import {moveToolbar} from './components/equipmentsMove/components/moveToolbar.js';

//компонент с вкладками
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
            header: "События выдачи",
            id: "regUserEvents",
            rows: [
                recordEventsDatepicker,
                recordEvent
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

//загрузка данных в древовидный список
function pushToTree(id) {
    let promise = "";
    let product = new Equipment();
    let token = getCurrentToken();
    $$(id).clearAll();
    promise = product.getTree(token);

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

//зазрука при переходе на панели, id - идентификатор каждой панели
function loadData(id) {
    if (id == "regProducts") {
        let promise = "";
        let equipment = new Equipment();
        let token = getCurrentToken();
        pushToTree(RegproductsTree);    //загрузка данных в древовидный список
        $$(TreeDatatable).clearAll();
        promise = equipment.getAll(token);  //получение всего оборудования

        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                if (result.Data != null) {
                    $$(TreeDatatable).parse(result.Data);
                    $$(TreeDatatable).filterByAll();
                } else {
                    $$(TreeDatatable).hideProgress({});
                }
            } else {
                webix.message(result.Error);
            }
        });
    }

    if (id == "regUsers") {
        $$(UsersDatatable).clearAll();
        let user = new Employee();
        let promise = user.getAll();
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                if (result.Data != null) {
                    $$(UsersDatatable).parse(result.Data);
                    $$(UsersDatatable).filterByAll();
                } else {
                    $$(UsersDatatable).hideProgress({});
                }
            } else {
                webix.message(result.Error);
            }
        })
    }

    if (id == "regUserEvents") {
        $$(EmployeeEventsDatatable).clearAll();
        let event = new UserEvent();
        let promise = event.getAll();
        promise.then(response => {
            return response.json();
        }).then(result => {
                if (result.Error == "") {
                    if (result.Data != null) {
                        $$(EmployeeEventsDatatable).clearAll();
                        $$(EmployeeEventsDatatable).parse(result.Data);
                        $$(EmployeeEventsDatatable).filterByAll();
                    } else {
                        $$(EmployeeEventsDatatable).hideProgress({});

                    }
                } else {
                    webix.message(result.Error);
                }
            }
        )
    }

    if (id == "moveProducts") {
        let promise = "";
        let product = new Equipment();
        pushToTree(MoveEquipmentTree);//parse Tree
        $$(MoveEquipmentTree).clearAll();
        $$(MoveEquipDatatable).clearAll();
        $$(DragProdDatatable).clearAll();

        promise = product.getEquipmentsInStore();
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                if (result.Data != null) {
                    $$(MoveEquipDatatable).parse(result.Data);
                    $$(MoveEquipDatatable).filterByAll();
                } else {
                    $$(MoveEquipDatatable).hideProgress({});
                }
            } else {
                webix.message(result.Error);
            }
        });
        //заполнение выпадающего списка в перемещении оборудования
        let user = new Employee();
        let users = user.getAll();

        users.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                if (result.Data != null) {
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
                }
            } else {
                webix.message(result.Error);
            }
        });
    }
}

//отображает все компоненты на странице
webix.ui({
    rows: [
        {view: "button", id: "authorize", value: "Выйти", width: 200, height: 50, align: "right", click: logout},
        mainPage,
    ]
});

//получение куки по имени
function getCookie(name) {
    let matches = document.cookie.match(new RegExp(
        "(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"
    ));
    return matches ? decodeURIComponent(matches[1]) : undefined;
}

//получение текущего токена
function getCurrentToken() {
    let cookie = getCookie("token");
    let splitCookie = cookie.split(':');
    return splitCookie[0];
}

//меняет отображение в зависимости от роли пользователя
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
    loadData("regProducts");
};

//выход из аккаунта
function logout() {
    let promise = fetch("/logout"); //запрос на выход
    promise.then(json => {
        return json.json()
    }).then(result => {
        if (result.Error == "") {
            document.cookie = "auth" + '=; Max-Age=0'; //удаление куки
            document.cookie = "token" + '=; Max-Age=0'; //удаление куки
            document.location.href = result.Data; //перенаправление по URI
        } else {
            webix.message(result.Error);
        }
    });
}

// добавление спиннеров
webix.extend($$(TreeDatatable), webix.ProgressBar);
webix.extend($$(UsersDatatable), webix.ProgressBar);
webix.extend($$(MoveEquipDatatable), webix.ProgressBar);
webix.extend($$(EmployeeEventsDatatable), webix.ProgressBar);
webix.extend($$(UsersDatatable), webix.ProgressBar);


//фильтр для для выборки элементов всех подклассов класса в древовидном списке учета обрудования
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

//фильтр для для выборки элементов подкласса в древовидном списке учета оборудования
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

//фильтр для для выборки элементов всех подклассов класса в древовидном списке перемещения оборудования
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

//фильтр для для выборки элементов подкласса в древовидном списке перемещения оборудования
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
    let product = new Equipment();
    let selected = $$(combo).getValue();

    if (selected == 2) {
        webix.message("Не выбран пользователь");
        return false;
    }
    if (value.Fk_user.Valid) { //проверка на то,в какую таблицу производится перемещение
        return false;
    }

    let sendValue = {fk_user: selected, id: value.id}; //ид пользователя и ид перемещаемого оборуования
    sendValue.fk_user = parseInt(sendValue.fk_user);
    let promise = product.dragToUser(sendValue);

    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            $$(DragProdDatatable).parse(result.Data);
            $$(MoveEquipDatatable).remove(result.Data.id);
        } else {
            webix.message(result.Error);
        }
    });
    return false;
});

//событие на drap-grop от сотрудника на склад
$$(MoveEquipDatatable).attachEvent("onBeforeDrop", function (context, ev) {
    let dnd = webix.DragControl.getContext();
    let value = dnd.from.getItem(dnd.source[0]);
    let product = new Equipment();

    if (!value.Fk_user.Valid) { //проверка на то,в какую таблицу производится перемещение
        return false;
    }
    let promise = product.dragToStore(value);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            $$(MoveEquipDatatable).parse(result.Data);
            $$(DragProdDatatable).remove(result.Data.id);
        } else {
            webix.message(result.Error);
        }
    });
    return false;
});

