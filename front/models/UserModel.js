export class User {
    constructor(user) {
        this.name = user.name;
        this.surname = user.surname;
        this.patronymic = user.patronymic;
        this.login = user.login;
        this.password = user.password;
    }
    updateUser() { }
    deleteUser() { }
    getUsers() { }
    resetPassword() { }
    authorize() { }
    registerUser() { }
}