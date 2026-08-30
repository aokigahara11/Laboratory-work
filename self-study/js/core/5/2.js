// Создать объект calculator с методами read(a, b), sum(), mul(). 
// Реализовать методы так, чтобы их можно было вызывать по цепочке (Chaining): calculator.read(2, 3).sum().

// Объект calculator
const calculator = {
    a: 0, 
    b: 0,

    // Методы: read, sum, mul
    read: function(a, b) {
        this.a = a;
        this.b = b;
        return this;
    },

    sum: function() {
        const result = this.a + this.b;
        return result
    },

    mul: function() {
        const result = this.a * this.b;
        return result;
    }
}

calculator.read(2, 3).sum()