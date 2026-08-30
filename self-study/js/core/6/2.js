// Создать класс Admin, наследующийся от User (extends), 
// добавляющий массив прав доступа permissions и метод grantPermission(perm).

class User {
    #name
    #password

    constructor(name, password) {
        this.name = name;
        this.#password = password;
    }
}

class Admin extends User {
    constructor(name, password) {
        super(name, password);
        this.permissions = []; 
    }

    grantPermission(perm) {
        this.permissions.push(perm);
        console.log(`Право "${perm}" успешно выдано!`);
    }
}