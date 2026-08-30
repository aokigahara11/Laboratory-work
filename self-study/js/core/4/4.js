// Написать функцию groupBy(array, property), которая группирует 
// массив объектов по заданному ключу с помощью .reduce() и возвращает объект групп.

function groupBy(array, property) {
    return array.reduce((acc, item) => {
        const key = item[property];
        
        if (!acc[key]) {
            acc[key] = [];
        }
        
        acc[key].push(item);
        
        return acc;
    }, {});
}
