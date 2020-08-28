import { UsersDatatable } from '../views/const.js';
export class User {

    updateUser(user) {
        return new Promise((resolve, object) => {
            if (true) {
                resolve(user);
            }
        });
    }
    deleteUser(user) {
        return new Promise((resolve, object) => {
            console.log(("userId",user.id));
            resolve(user.id);
        });
    }
    resetPassword(user) {
        return new Promise((resolve, reject) => {
            if (true) {
                resolve(user);
            }
        });
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

        this.id = +($$("usersList").getLastId()) + 1;
        console.log(this.id);
        return new Promise((resolve, reject) => {
            if (true) {
                resolve(this);
            }
        });
    }
    getAllUsers(user,id) {
        if (id != "") {
            $$(UsersDatatable).showProgress({});
        }
        return new Promise((resolve, reject) => {
            let arr = [
                { id: 1, name: "Ivan", surname: "Ivanovich", patronymic: "Ivanov", login: "IvanIvan", },
                { id: 233, name: "Ivan", surname: "Ivanovich", patronymic: "Ivanov", login: "IvanIvan", },
                { id: 11, name: "Ivan", surname: "Ivanovich", patronymic: "Ivanov", login: "IvanIvan", }
            ];
            setTimeout(() => {
                resolve(arr);
            }, 0);
        })
    }

}