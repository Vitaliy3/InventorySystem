import { toolbar,toolbar_user } from './dependViews/toolBar.js';
import { treeArrayFull } from './dependViews/regProductView.js';
import { treeArrayFull_user } from './dependViews/regProductView_user.js';



const mainPage = {
    width: 200,
    header: "TESTING",
    height: 1000,
    view: "tabview",
    cells: [
        {
            header: "Учет оборудования",

            rows: [
                /*  {
                      cols: [
                          { view: "text", id: "filterClass", width: 100 },//filter
                        
                      ]
                  },*/
                toolbar,
                treeArrayFull,

            ]
        },
        {
            header: "Учет оборудования(Сотрудник)",

            rows: [
                treeArrayFull_user,

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
        mainPage
    ]
});

