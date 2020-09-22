import {sendQuery} from "./MEquipmentM.js";
import {UsersDatatable} from "../views/const.js";

export class Employee {
    update(user) {
        let json = JSON.stringify(user);
        console.log(json);
        return sendQuery('/updateEmployee', 'POST', json);
    }

    delete(user) {
        let json=JSON.stringify(user);
        console.log(json);
        return sendQuery('/deleteEmployee', 'DELETE', json);
    }

    resetPassword(user) {
        let json=JSON.stringify(user);
        return sendQuery('/resetPassEmployee', 'POST',json);
    }

    register(user) {
        let json = JSON.stringify(user);
        console.log(json);

        return sendQuery('/addEmployee', 'POST', json);

    }

    getAll() {
        $$(UsersDatatable).showProgress({});
        return fetch('/getAllEmployees')
    }
}
