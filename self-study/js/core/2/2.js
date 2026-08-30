// Задание 2.2: Реализовать программу «FizzBuzz» с помощью цикла for от 1 до 100: для кратных 3 выводить "Fizz", 
// для кратных 5 — "Buzz", для кратных и 3, и 5 — "FizzBuzz".

function FizzBuzz() {
    for (let i = 1; i <= 100; i++) {
        if (i % 3 === 0) {
            console.log(`Fizz`);
        } else if (i % 5 === 0) {
            console.log(`Buzz`);
        } else if (i % 5 === 0 && i % 3 === 0) {
            console.log(`FizzBuzz`);
        }
    }
}