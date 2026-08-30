// Проверить работу строгого (===) и нестрогого (==) сравнения 
// для пар: 5 и "5", null и undefined, 0 и false. Вывести результаты в консоль с пояснением.

function comparison(a, b) {
    console.log(`Результат сравнения: нестрогое ${a == b}, строгое ${a === b}.`);
}


function main() {
    comparison(5, "5");
    comparison(null, undefined);
    comparison(0, false);
    console.log(`Не строгое сравнение сравнивает только значение, а строгое сравнивает значение и тип.`)
}

