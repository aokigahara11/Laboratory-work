// Дан массив объектов пользователей [{id: 1, name: "A", age: 25}, ...]. 
// С помощью метода .sort() отсортировать их по возрасту по убыванию.

const users = [
    { id: 1, name: "A", age: 25 },
    { id: 2, name: "B", age: 30 },
    { id: 3, name: "C", age: 22 },
    { id: 4, name: "D", age: 35 },
    { id: 5, name: "E", age: 28 }
];

function ageSort(usersData) {
    usersData.sort((a, b) => {
        if (a.age < b.age) {
            return 1;
        } else if (a.age > b.age) {
            return -1;
        } else {
            return 0;
        }
    });
}