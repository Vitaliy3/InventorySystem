import { UsersList } from '../views/const.js';
export class User {
    constructor(user) {
        this.id = user.id;
        this.name = user.name;
        this.surname = user.surname;
        this.patronymic = user.patronymic;
        this.login = user.login;
        this.password = user.password;
    }
    updateUser() {
        return new Promise((resolve, object) => {
            if (true) {
                resolve(this);
            }
        });
    }
    deleteUser() {
        return new Promise((resolve, object) => {
            resolve(this);
        });
    }
    resetPassword() {
        return new Promise((resolve, reject) => {
            if (true) {
                resolve(this);
            }
        });
    }
    authorize() {
        return new Promise((resolve, reject) => {
            console.log(this);
            if (this.login == "1" && this.password == "1") {
                resolve("Admin");
            } else {
                resolve("Employee");
            }
        });
    }
    registerUser() {

        this.id = +($$("usersList").getLastId()) + 1;
        console.log(this.id);
        return new Promise((resolve, reject) => {
            if (true) {
                resolve(this);
            }
        });
    }
    getAllUsers(id) {
        if (id != "") {
            $$(UsersList).showProgress({});
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