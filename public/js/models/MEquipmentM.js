import {MoveEquipDatatable, TreeDatatable} from '../views/const.js';

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
        let json = JSON.stringify(equipment);
        return sendQuery('/deleteEquipment', 'DELETE', json)
    }

    updateEquipment(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/updateEquipment', 'POST', json)
    }

    getAllEquipments(token) {
        $$(TreeDatatable).showProgress({});
        return fetch('/getAllEquipments?token=' + token)
    }

    getEquipmentsInStore() {
        $$(MoveEquipDatatable).showProgress({});
        return fetch('/getEquipmentsInStore')
    }

    getEquipmentsByUser(selectedEmployee) {
        return fetch('/getEquipmentByUser?user=' + selectedEmployee);
    }

    getFullTree(token) {
        return fetch('/getFullTree?token=' + token);
    }


    writeEquipment(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/writeEquipment', 'POST', json);
    }

    dragToUser(equipment) {
        console.log(equipment);

        let unParsed = JSON.stringify(equipment);
        return sendQuery('/dragToUser', 'POST', unParsed);
    }

    dragToStore(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/dragToStore', 'POST', json);
    }
}

