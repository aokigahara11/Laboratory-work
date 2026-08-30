// Задание 2.3: Написать функцию isPalindrome(str), которая принимает строку, 
// приводит ее к нижнему регистру, удаляет пробелы и проверяет, является ли она палиндромом.

function isPalindrome(str) {
    str = str.toLowerCase();
    str = str.trim();

    let len = str.length()
    for (let i = 0; i < len / 2; i++) {
        if (str[i] != str[i+1]) {
            return false
        }
    } 
    return true    
}