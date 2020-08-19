var selected = treeElements;
var treeClass = [
    {
        id: "1", value: "Oceansize", data: [
            { id: "treeElementsArr1", value: "Everyone Into Position" },
        ]
    },
    {
        id: "2", value: "Little People", data: [
            { id: "treeElementsArr2", value: "Mickey Mouse Operation" },
        ]
    },
];
var treeElementsArr1 = [
    { id: "1", name: "01. Basique", user: "3:38", status: "1", inventoryNumber: "11" },
];
var treeElementsArr2 = [
    { id: "1", title: "01. Basique", duration: "3:38" },
    { id: "2", title: "02. Moon", duration: "3:47" },
    { id: "3", title: "03. Unsaid", duration: "3:48" },
    { id: "4", title: "02. Moon", duration: "3:47" },
    { id: "5", title: "03. Unsaid", duration: "3:48" }
];
var gridColumns = [
    {
        dataIndex: "title",
        header: "Title"
    },
    {
        dataIndex: "duration",
        header: "Duration"
    }
];

var treeElements = [
    { id: "1", title: "01. The Charm Offensive", duration: "7:19" },
    { id: "2", title: "02. Heaven Alive", duration: "6:20" },
    { id: "3", title: "03. A Homage to Shame", duration: "5:52" },

];
export const treeArrayFull = {
    cols: [
        {
            rows: [
                {
                    view: "tree", id: "myTree", width: 250, data: treeClass, select: treeElements, on: {
                        onSelectChange: function () {
                            selected = $$("myTree").getSelectedId();
                            if (isNaN(selected)) {
                                $$("myList").clearAll();
                                $$("myList").parse(window[selected]); // reference to the id-matching variable
                            }
                        }
                    }
                },

            ]
        },
        { view: "resizer" },
        {
            rows: [
                {
                    view: "datatable",
                    id: "myList",
                    autoConfig: true,
                    data: selected,
                    columns: [
                        { id: "name", header: ["Название", { content: "selectFilter" }], width: 200, },
                        { id: "user", header: ["Сотрудник", { content: "selectFilter" }], width: 200 },
                        { id: "status", header: ["Статус", { content: "selectFilter" }], width: 100 },
                        { id: "inventoryNumber", header: ["Инвентарный номер", { content: "selectFilter" }], width: 200 },
                    ]
                }]
        }
    ]
};