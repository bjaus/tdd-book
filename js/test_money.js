const assert = require('assert');

class Money {
  constructor(amount, currency) {
    this.amount = amount;
    this.currency = currency;
  }

  times(multipler) {
    return new Money(this.amount * multipler, this.currency)
  }

  divide(divisor) {
    return new Money(this.amount / divisor, this.currency);
  }
}

let money = new Money(5, "USD");
let expected = new Money(10, "USD")
assert.deepStrictEqual(money.times(2), expected);

money = new Money(10, "EUR")
expected = new Money(20, "EUR")
assert.deepStrictEqual(money.times(2), expected);

money = new Money(4002, "KRW")
expected = new Money(1000.5, "KRW")
assert.deepStrictEqual(money.divide(4), expected)
