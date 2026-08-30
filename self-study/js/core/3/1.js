// Написать функцию calculateTotal(price, discount = 0, tax = 0.2), 
// использующую параметры по умолчанию и возвращающую итоговую сумму.

function calculateTotal(price, discount = 0, tax = 0.2) {
    let discountedPrice = price * (1 - discount);
    let total = discountedPrice * (1 + tax);
    return total;
}