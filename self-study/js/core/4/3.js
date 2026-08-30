// Используя метод .reduce(), посчитать общую стоимость 
// товаров в корзине покупок: [{name: "Laptop", price: 1000, count: 2}, ...].

function getPrices(data) {
    const totalCost = data.reduce((acc, item) => {
    if (item.count > 0) {
        return acc + (item.price * item.count);
    } else {
        return acc;
    }
}, 0);

return totalCost
}

