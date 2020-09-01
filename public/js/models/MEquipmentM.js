import {TreeDatatable} from '../views/const.js';

export function sendQuery(url, method, data) {
    let response = fetch(url, {
        method: method,
        headers: {
            'Content-Type': 'application/json'
        },
        body: data
    });
    return response;
}

export class Equipment {
    addEquipment(equipment) {
        equipment.class = parseInt(equipment.class);
        equipment.subclass = parseInt(equipment.subclass);
        let json = JSON.stringify(equipment);
        return sendQuery('/addEquipment', 'POST', json)
    }

    deleteEquipment(equipment) {
        return sendQuery('/deleteEquipment', 'DELETE', equipment.id)
    }

    updateEquipment(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/updateEquipment', 'POST', json)
    }

    getAllEquipment() {
        return fetch('/getAllEquipments')
    }


    getEquipmentsInStore(product) {
        $$(TreeDatatable).showProgress({});
        return new Promise((resolve, object) => {
            let arr = [
                {
                    id: "1",
                    class: "1",
                    subclass: "1",
                    name: "name1",
                    user: "user1",
                    status: "на складе",
                    inventoryNumber: "11"
                },
                {
                    id: "2",
                    class: "1",
                    subclass: "1",
                    name: "name1",
                    user: "user1",
                    status: "на складе",
                    inventoryNumber: "11"
                },
                {
                    id: "3",
                    class: "1",
                    subclass: "1",
                    name: "name1",
                    user: "user1",
                    status: "на складе",
                    inventoryNumber: "11"
                },
                {
                    id: "4",
                    class: "1",
                    subclass: "1",
                    name: "name1",
                    user: "user1",
                    status: "на складе",
                    inventoryNumber: "11"
                },
                {
                    id: "5",
                    class: "1",
                    subclass: "1",
                    name: "name1",
                    user: "user1",
                    status: "на складе",
                    inventoryNumber: "11"
                },
                {
                    id: "6",
                    class: "2",
                    subclass: "2",
                    name: "name2",
                    user: "user2",
                    status: "на складе",
                    inventoryNumber: "11"
                },
                {
                    id: "7",
                    class: "2",
                    subclass: "2",
                    name: "name2",
                    user: "user2",
                    status: "на складе",
                    inventoryNumber: "11"
                },
            ];
            setTimeout(() => {
                resolve(arr);
            }, 10);
        });
    }

    getUserProducts(product) {
        fetch('http://192.168.77.142:9000/getEquipment?user="taps"').then(res => {
            return res.json()
        }).then(res => {
        });
    }

    getUserClasses(user) {
        return new Promise((resolve, object) => {
            let str = [
                {
                    class: "0", value: "All", open: true, data: [
                        {
                            class: "1", value: "Столы", open: true, data: [
                                {subclass: "1", value: "Компьютерный"},
                            ]
                        },
                    ]
                }];

            resolve(str);
        });
    }

    getAllTree() {
        return fetch('/getFullTree')
    }

    writeProduct(equipment) {
        return sendQuery('/writeEquipment', 'POST', equipment.id);
    }

    moveProduct(product) {
    }

    dragToUser() {
        return true;
    }

    dragToStore() {
        return true;
    }
}

