
import { toolbar } from './dependViews/toolBar.js';
import { tree } from './dependViews/tree.js';
import { registerUsers } from './dependViews/registerUsersView.js';


const mainPage = {
    width: 200,
    header: "TESTING",
    height: 1000,
    id: "tabView",
    view: "tabview",
    cells: [
        {
            header: "Учет оборудования",

            rows: [
                toolbar,
                tree,

            ]
        },
        {
            header: "Учет сотрудников",
            rows: [
                registerUsers,
            ]
        },
        {
            rows: [
                //next tab,
            ]
        },

    ]
};




webix.ui({
    rows: [
        mainPage,

    ]
});