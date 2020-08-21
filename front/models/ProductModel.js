export class Product {
    constructor(product) {
        this.id = product.id;
        this.name = product.name;
        this.user = product.user;
        this.status = product.status;
        this.inventoryNumber = product.inventoryNumber;
        this.class=this.class;
        this.subclass=this.subclass;
    }
    addProduct() {
        return new Promise((resolve, reject) => {
            this.id = +($$('myList').getLastId()) + 1;
            if (true) {
                resolve(this);
            } else {
                let error = new Error(get.status);
                error.code = get.status;
                reject(error);
            }
        });
    }

    deleteProduct() {
        return new Promise((resolve, object) => {
            resolve(this);
        });
    }
    updateProduct() {
        return new Promise((resolve, object) => {
            if (true) {
                resolve(this);
            } else {
                let error = new Error(get.status);
                error.code = get.status;
                reject(error);
            }
        });
    }
    getAllProducts() {
        return [
            { id: "1", fkClass: "1", fk_subClass: "1", name: "name1", user: "user1", status: "ok", inventoryNumber: "11" },
            { id: "2", fkClass: "1", fk_subClass: "1", name: "name1", user: "user1", status: "ok", inventoryNumber: "11" },
            { id: "3", fkClass: "1", fk_subClass: "1", name: "name1", user: "user1", status: "ok", inventoryNumber: "11" },
            { id: "4", fkClass: "1", fk_subClass: "1", name: "name1", user: "user1", status: "ok", inventoryNumber: "11" },
            { id: "5", fkClass: "1", fk_subClass: "1", name: "name1", user: "user1", status: "ok", inventoryNumber: "11" },
            { id: "6", fkClass: "2", fk_subClass: "2", name: "name2", user: "user2", status: "ok", inventoryNumber: "11" },
            { id: "7", fkClass: "2", fk_subClass: "1", name: "name2", user: "user1", status: "ok", inventoryNumber: "11" },
        ];
    }
    pushClassSub() {

    }
    pushProducts() {

    }

    getClassSubclass() {
        let str = [
            {
                class: "1", value: "Столы", data: [
                    { subclass: "1", value: "Компьютерный" },
                ]
            },
            {
                class: "2", value: "Стулья", data: [
                    { subclass: "1", value: "Для офиса" },
                    { subclass: "2", value: "Для дома" },
                ]
            },
        ];


        return str;
    }

    writeProduct() {
        return new Promise((resolve, object) => {
            this.status = "Writed";
            resolve(this);
        });
    }
    moveProduct() {
    }
}