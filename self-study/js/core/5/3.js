// 3: Продемонстрировать потерю контекста this при передаче метода объекта в setTimeout 
// и исправить проблему тремя способами: через обертку-функцию, метод .bind() и стрелочную функцию.

const user = {
  username: "Иван",
  
  PrintHi: function() {
    console.log(`Привет, ${this.username}!`);
  }
};

// setTimeout вызывает PrintHi без объекта
setTimeout(user.PrintHi, 1000); 

// 1. Обертка-функция
setTimeout(function() {
  user.PrintHi(); // Вызов внутри идет через точку — контекст сохранен
}, 1000);

// 2. Метод .bind()
// Создает новую функцию с жестко привязанным this = user
setTimeout(user.PrintHi.bind(user), 1000);

// 3. Стрелочная функция
// Берет user из замыкания и вызывает метод через точку
setTimeout(() => user.PrintHi(), 1000);
