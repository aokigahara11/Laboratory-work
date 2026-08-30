// Написать функцию-счетчик createCounter(initialValue), 
// которая использует замыкание (Closure) и возвращает объект с методами increment(), decrement(), getValue(), reset().

function createCounter(initialValue) {
    let count = initialValue

    // Замыкание
    return {
        increment() {
            count++;
        },
        decrement() {
            count--;
        },
        getValue() {
            return count;
        },
        reset() {
            count = initialValue;
            return count;
        }
    };
}