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
    add(equipment) {
        equipment.class = parseInt(equipment.class);
        equipment.subclass = parseInt(equipment.subclass);
        let json = JSON.stringify(equipment);

        return sendQuery('/addEquipment', 'POST', json)
    }

    delete(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/deleteEquipment', 'DELETE', json)
    }

    update(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/updateEquipment', 'POST', json)
    }

    getAll(token) {
        $$(TreeDatatable).showProgress({});
        return fetch('/getAllEquipments?token=' + token)
    }

    getEquipmentsInStore() {
        $$(MoveEquipDatatable).showProgress({});
        return fetch('/getEquipmentsInStore')
    }

    getByUser(selectedEmployee) {
        return fetch('/getEquipmentByUser?user=' + selectedEmployee);
    }

    getTree(token) {
        return fetch('/getFullTree?token=' + token);
    }


    write(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/writeEquipment', 'POST', json);
    }

    dragToUser(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/dragToUser', 'POST', json);
    }

    dragToStore(equipment) {
        let json = JSON.stringify(equipment);
        return sendQuery('/dragToStore', 'POST', json);
    }
}

