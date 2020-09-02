import {equipmentsToolbar} from './dependViews/forms/equipmentsToolbar.js';
import {hide, recordEquipments} from './dependViews/recordEquipments.js';
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
import {EmployeeEvent} from '../models/MEmployeeEvents.js';
import {moveToolbar, movingTree} from './dependViews/equipmentsMove.js';
import {authorizeForm} from './dependViews/forms/authorization.js';

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
                recordEventsDatapicker,
                recordEvents
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
                if (result.Error == "") {
                    console.log("tree", result.Tree);
                    $$(id).parse(result.Tree);
                } else {
                    webix.message(result.Error);
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
            promise = product.getUserEquipments();
        }
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(TreeDatatable).parse(result.DataArray);
                $$(TreeDatatable).filterByAll();
            } else {
                webix.message("err", result.Error);
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
        let event = new EmployeeEvent({});
        let promise = event.getAllEvents();
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(UserEventsDatatable).clearAll();
                $$(UserEventsDatatable).parse(result.DataArray);

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
        if (!hide) {
            let product = new Equipment({});
            promise = product.getEquipmentsInStore();
        }
        promise.then(response => {
            return response.json();
        }).then(result => {
            if (result.Error == "") {
                $$(MoveEquipDatatable).parse(result.DataArray);
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
                for (let i = 0; i < result.DataArray.length; i++) {
                    temp = {
                        id: result.DataArray[i].id,
                        name: result.DataArray[i].name + " " + result.DataArray[i].surname + " " + result.DataArray[i].patronymic
                    };
                    joinUsers.push(temp);
                }
                let list = $$(combo).getPopup().getList();
                list.clearAll();
                list.parse(result.DataArray);
                console.log(joinUsers);
            } else {
                webix.message(result.Error);
            }
        });
    }
}

webix.ui({
    rows: [
        {view: "button", id: "authorize", value: "Войти", width: 200, height: 50, align: "right", click: authorize},//временно расположена здесь
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
    sendValue.fk_user=parseInt(sendValue.fk_user);
    let promise = product.dragToUser(sendValue);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            return result.Data;
        } else {
            webix.message(result.Error);
            return false;
        }
    })
    return false;
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
            return result.Data;
        } else {
            webix.message(result.Error);
        }
    });
    ;
    return false;
});
