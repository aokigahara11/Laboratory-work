// Задание 4.1: Дан массив чисел. С помощью метода .filter() 
// отфильтровать только нечетные числа, а с помощью .map() возвести их в квадрат.

function Filter(numbers) {
    const result = numbers.filter(num => num % 2 == 0).map(num => num ** 2)
    return result
}
