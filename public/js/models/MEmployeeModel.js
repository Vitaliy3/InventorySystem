import {sendQuery} from "./MEquipmentM.js";
import {UsersDatatable} from "../views/const.js";

export class Employee {

    updateUser(user) {
        let json = JSON.stringify(user);
        console.log(json);
        return sendQuery('/updateEmployee', 'POST', json);

    }

    deleteUser(user) {
        let json=JSON.stringify(user);
        console.log(json);
        return sendQuery('/deleteEmployee', 'DELETE', json);
    }

    resetPassword(user) {
        let json=JSON.stringify(user);

        return sendQuery('/resetPassEmployee', 'POST',json);
    }

    authorize(user) {
        return new Promise((resolve, reject) => {
            console.log(this);
            if (this.login == "1" && this.password == "1") {
                resolve("Admin");
            } else {
                resolve("Employee");
            }
        });
    }

    registerUser(user) {
        let json = JSON.stringify(user);
        console.log(json);
        return sendQuery('/addEmployee', 'POST', json);

    }

    getAllEmployees() {
        $$(UsersDatatable).showProgress({});
        return fetch('/getAllEmployees')
    }
}