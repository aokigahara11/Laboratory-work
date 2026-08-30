// Продемонстрировать работу синтаксиса Destructuring (деструктуризация массива и объекта) 
// и оператора Nullish Coalescing (??) на примере парсинга объекта настроек пользователя.

const userData = {
    id: 1,
    name: 'Иван',
    settings: {
        theme: null,
        language: 'ru',
        notifications: {
            email: false
        }
    }
};

function getUserData(userData) {
    const {
        id,
        name,
        settings: {
            theme = 'light',
            language = 'en',
            notifications: {
                email = true,
                push = true
            } = {}
        } = {}
    } = data;

    return {
        id,
        name,
        theme: theme ?? 'light',
        language: language ?? 'en',
        notifications: {
            email: email ?? false,
            push: push ?? true
        }
    };
}


