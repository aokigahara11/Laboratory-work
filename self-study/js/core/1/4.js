// Написать проверку типа переменной через typeof. 
// Обработать случаи для числа, строки, boolean, undefined, null (учесть баг typeof null === "object"), массива и объекта.

function checkType(a) {
    if (a === null) {
        return "null";
    }
    
    switch (typeof a) {
        case "number":
            return "number";
        case "string":
            return "string";
        case "boolean":
            return "boolean";
        case "undefined":
            return "undefined";
        case "object":
            return "object";
        default:
            return typeof a;
    }
}