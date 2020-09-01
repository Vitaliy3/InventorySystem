import {sendQuery} from "./MEquipmentM.js";

export class Employee {

    updateUser(user) {
        let json = JSON.stringify(user);
        console.log(json);
        return sendQuery('/updateEmployee', 'POST', json);

    }

    deleteEmployee(user) {
        return sendQuery('/deleteEmployee', 'DELETE', user.id);
    }

    resetPassword(user) {
        return sendQuery('/resetPassEmployee', 'POST', user.id);
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
        return fetch('/getAllEmployees')

    }

}