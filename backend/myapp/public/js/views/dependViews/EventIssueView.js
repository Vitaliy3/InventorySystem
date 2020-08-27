 export const regEventIssue = {
    header: "Учет событий выдачи",
    body: {
        view: "datatable",
        columns: [
            { id: "name", header: "name", width: 100, footer: { text: "summa", colspan: 2 } },
            { id: "count", header: "count", width: 100 },
            { id: "price", header: "price", width: 100, footer: { content: "summColumn" } },
        ],
        on: {
            onItemClick: function () {
                console.log('ok');
            }
        },
        type: {
            height: 60
        },
        select: true,
        data: big_film_set
    },
};