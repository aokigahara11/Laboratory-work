// Написать функцию deepClone(obj), 
// которая делает глубокую копию объекта без использования 
// JSON.parse(JSON.stringify(obj)) и библиотеки Lodash.

function deepClone(obj) {
    if (obj === null || typeof obj !== 'object') {
        return obj;
    }

    const copy = Array.isArray(obj) ? {} : []; // обьект или массив

    for (let key in obj) {
        if (Object.prototype.hasOwnProperty.call(obj, key)) {
        copy[key] = deepClone(obj[key]); // Рекурсивно вызываем функцию для каждого вложенного значения
        }
    }

  return copy;
}