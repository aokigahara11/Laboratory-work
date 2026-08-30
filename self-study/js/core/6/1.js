// Создать класс User с приватным полем #password, геттером/сеттером для имени и методом checkPassword(pass).

class User {
    #password;
    #name;

    constructor(name, password) {
        this.name = name;
        this.#password = password;
    }

    get name() {
        return this.#name;
    }

    set name(newName) {
        if (!newName || newName.trim() === "") {
            console.log("Ошибка: имя не может быть пустым!");
            return;
        }
        this.#name = newName.trim();
    }

    checkPassword(pass) {
        return this.#password === pass;
    }
}
