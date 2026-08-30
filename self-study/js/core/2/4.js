// Задание 2.4: Написать функцию truncate(str, maxlength), 
// которая проверяет длину строки str и, если она превышает maxlength, 
// заменяет конец строки на "...".

function truncate(str, maxlength) {
    let len = str.length()

    for (let i = 0; i < len; i++) {
        if (i > maxlength) {
            str.replace(str[i], "...")
        }
    }
}