import {UserEventsDatatable} from '../views/const.js';

export class UserEvent {
    getAllEvents() {
        $$(UserEventsDatatable).showProgress({});
        return fetch('/getAllEvents')
    }


    getEventsForDate(dateFromTo) {
        let date = JSON.parse(dateFromTo);
        console.log(date);
        date.start = date.start.split(' ')[0];
        date.end = date.end.split(' ')[0];
        return fetch('/getEventsForDate/?dateStart=' + date.start + '&dateEnd=' + date.end);
    }
}