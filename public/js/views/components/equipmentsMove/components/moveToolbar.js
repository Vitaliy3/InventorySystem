import {combo, DragProdDatatable} from "../../../const.js";
import {Equipment} from "../../../../models/MEquipmentM.js";

//панель инстурментов для панели учета оборудования
export const moveToolbar = {
    view: "toolbar",
    id: "moveToolbar",
    width: 200,
    cols: [
        {
            view: "template",
            css: {"opacity": "0"},
            template: "<div></div>"
        },
        {
            view: combo,
            value: 2,
            id: combo,
            width: 295,
            align: "right",
            options: {
                body: {
                    template: "<span style='display:none'>#id#</span> <span>#name#</span>"
                },
            }
        },
        {view: "button", id: "button1", value: "Найти", click: findEquipmentsByUser, width: 100,},
    ],
};

//поиск оборудования по выбранному пользователю
function findEquipmentsByUser() {
    let eqipment = new Equipment();
    let selected = $$(combo).getValue();
    $$(DragProdDatatable).clearAll();

    let promise = eqipment.getByUser(selected);
    promise.then(response => {
        return response.json();
    }).then(result => {
        if (result.Error == "") {
            if (result.Data == null) {
            } else {
                console.log("data", result.Data);
                $$(DragProdDatatable).parse(result.Data);
            }
        } else {
            webix.message(result.Error);
        }
    });
}