// Объявить переменные firstName, lastName, age с помощью const и let. Сформировать строку приветствия с помощью шаблонных строк (`) вида: 
// "Привет, меня зовут [firstName] [lastName], мне [age] лет".

function PrintInfo() {
    const firstName = "Yatoro";
    const lastName = "God";
    let age = 16;

    console.log(`Привет, меня зовут ${firstName} ${lastName}, мне ${age} лет.`)
}