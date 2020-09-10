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

    getAllEquipment(token) {
        return fetch('/getAllEquipments?token='+token)
    }

    getEquipmentsInStore() {
        $$(TreeDatatable).showProgress({});
        return fetch('/getEquipmentsInStore')
    }

    getUserEquipments(selectedEmployee) {
        return fetch('/getEquipmentByUser?user=' + selectedEmployee);
    }

    getFullTree(token) {
        return fetch('/getFullTree?token=' + token);
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

