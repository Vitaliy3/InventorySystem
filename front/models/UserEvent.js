
import { UserEvent } from '../views/const.js';
export class InventoryEvent {
    constructor(event) {
        this.description = event.description;
        this.dateFrom = event.dateFrom;
        this.dateTo = event.dateTo;
        this.product = event.product;

    }
    addEvent() {

    }
    getEventsDate() {
        return new Promise((resolve, object) => {
            let events = [
                { id: "1", user: "Ivan", date: "01.01.2020", event: "Выдан сотруднику", product: "Стол ***" },
                { id: "2", user: "Ivan1", date: "01.01.2020", event: "Выдан сотруднику", product: "Стол ***" },
                { id: "3", user: "Ivan2", date: "01.01.2020", event: "Выдан сотруднику", product: "Стол ***" },
            ];
            resolve(events);
        });
    }
    getAllEvents() {
        $$(UserEvent).showProgress({
            hide: true
        });
        return new Promise((resolve, object) => {
            let date = "2020.08.01";
            let date1 = "2020.08.07";
            let date2 = "2020.08.11";
            let allEvents = [
                { id: "1", user: "Ivan", date: date, event: "Выдан сотруднику", product: "Стол ***" },
                { id: "2", user: "Ivan1", date: date1, event: "Выдан сотруднику", product: "Стол ***" },
                { id: "3", user: "Ivan2", date: date2, event: "Выдан сотруднику", product: "Стол ***" },
                { id: "4", user: "Ivan", date: date1, event: "Выдан сотруднику", product: "Стол ***" },
                { id: "5", user: "Ivan", date: date, event: "Выдан сотруднику", product: "Стол ***" },
                { id: "6", user: "Ivan3", date: date2, event: "Выдан сотруднику", product: "Стол ***" },
                { id: "7", user: "Ivan4", date: date, event: "Возвращен на склад", product: "Стол ***" },
                { id: "8", user: "Ivan", date: date1, event: "Возвращен на склад", product: "Стол ***" },
            ];
            setTimeout(() => {
                resolve(allEvents);
            }, 500);
        });

    }
    getUserEvents() {

    }

}


