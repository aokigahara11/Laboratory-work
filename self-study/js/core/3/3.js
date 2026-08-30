// Написать функцию mergeArrays(...arrays), которая принимает произвольное количество массивов 
// через Rest-оператор (...) и объединяет их в один плоский массив.

function mergeArrays(...arrays) {
    const count = arrays.length()
    let newArray = []
    return newArray.concat(...arrays)
}