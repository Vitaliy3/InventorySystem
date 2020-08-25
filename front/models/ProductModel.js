import { TreeList } from '../views/const.js';
export class Product {
    constructor(product) {
        this.id = product.id;
        this.name = product.name;
        this.user = product.user;
        this.status = product.status;
        this.inventoryNumber = product.inventoryNumber;
        this.class = this.class;
        this.subclass = this.subclass;
    }
    addProduct() {
        return new Promise((resolve, reject) => {
            this.id = +($$('myList').getLastId()) + 3;
            if (true) {
                resolve({
                    id: this.id, class: this.class, subclass: this.subclass, name: this.name, status: "На складе", user: "Ivan", inventoryNumber: "11"
                });
            }
        });
    }

    deleteProduct() {
        return new Promise((resolve, object) => {
            resolve(this);
        });
    }
    updateProduct() {
        //console.log(this);

        return new Promise((resolve, object) => {
            if (true) {
                resolve(this);
            }
        });
    }
    getAllProducts() {
        $$(TreeList).showProgress({});
        return new Promise((resolve, object) => {
            let arr = [
                { id: "1", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "2", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "3", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "4", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "5", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "6", class: "2", subclass: "2", name: "name2", user: "user2", status: "на складе", inventoryNumber: "11" },
                { id: "7", class: "2", subclass: "2", name: "name2", user: "user2", status: "на складе", inventoryNumber: "11" },
            ];
            setTimeout(() => {
                resolve(arr);
            }, 10);
        });
    }
    getProdutsInStore() {
        $$(TreeList).showProgress({});
        return new Promise((resolve, object) => {
            let arr = [
                { id: "1", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "2", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "3", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "4", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "5", class: "1", subclass: "1", name: "name1", user: "user1", status: "на складе", inventoryNumber: "11" },
                { id: "6", class: "2", subclass: "2", name: "name2", user: "user2", status: "на складе", inventoryNumber: "11" },
                { id: "7", class: "2", subclass: "2", name: "name2", user: "user2", status: "на складе", inventoryNumber: "11" },
            ];
            setTimeout(() => {
                resolve(arr);
            }, 10);
        });
    }
    getProductsUser() {
        return new Promise((resolve, object) => {
            let arr = [
                { id: "11", class: "1", subclass: "1", name: "name5", user: "User", status: "на складе", inventoryNumber: "11" },
                { id: "22", class: "1", subclass: "1", name: "name5", user: "User", status: "на складе", inventoryNumber: "11" },
                { id: "33", class: "1", subclass: "1", name: "name5", user: "User", status: "на складе", inventoryNumber: "11" },
            ];
            setTimeout(() => {
                resolve(arr);
            }, 10);
        });
    }
    getUserClasses() {
        return new Promise((resolve, object) => {
            let str = [
                {
                    class: "0", value: "All", open: true, data: [
                        {
                            class: "1", value: "Столы", open: true, data: [
                                { subclass: "1", value: "Компьютерный" },
                            ]
                        },
                    ]
                }];

            resolve(str);
        });
    }
    getAllClasses() {
        return new Promise((resolve, object) => {
            let str = [
                {
                    class: "0", value: "All", open: true, data: [
                        {
                            class: "1", value: "Столы", open: true, data: [
                                { subclass: "1", value: "Компьютерный" },
                            ]
                        },
                        {
                            class: "2", value: "Стулья", open: true, data: [
                                { subclass: "2", value: "Для офиса" },
                                { subclass: "3", value: "Для дома" },
                            ]
                        },

                    ]
                }];
            resolve(str);
        });
    }

    writeProduct() {
        return new Promise((resolve, object) => {
            this.status = "Writed";
            resolve(this);
        });
    }
    moveProduct() {
    }
    dragToUser() {
        return true;
    }
    dragToStore() {
        return true;
    }
}

