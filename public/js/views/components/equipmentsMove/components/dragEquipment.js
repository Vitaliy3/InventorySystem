import {DragProdDatatable} from "../../../const.js";

//представление: оборудование по выбранному сотруднику
export const dragEquipmentTable = {
    view: "datatable",
    drag: true,
    id: DragProdDatatable,
    width: 400,
    select: true,
    columns: [
        {id: "Class", header: "Класс", fillspace: true,},
        {id: "Subclass", header: "Подкласс", width:150,},
        {id: "name", header: "Название", fillspace: true,},
        {id: "inventoryNumber", header: "Инвентарный номер", fillspace: true},
    ]
};