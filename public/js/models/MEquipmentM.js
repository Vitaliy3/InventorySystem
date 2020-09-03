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

    getEquipmentsInStore() {
        $$(TreeDatatable).showProgress({});
        return fetch('/getEquipmentsInStore')
    }

    getUserEquipments(selectedEmployee) {
        return fetch('/getEquipmentOnUser?user=' + selectedEmployee);
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

    dragToUser(equipment) {
        let unParsed = JSON.stringify(equipment);
        console.log(unParsed);
        return sendQuery('/dragToUser', 'POST', unParsed);
    }

    dragToStore(equipment) {
        return sendQuery('/dragToStore', 'POST', equipment.id);
    }
}

