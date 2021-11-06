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

class Portfolio {
  constructor() {
    this.moneys = [];
  }

  add(...moneys) {
    this.moneys = this.moneys.concat(moneys);
  }

  evaluate(currency) {
    let total = this.moneys.reduce((sum, money) => {
      return sum + money.amount;
    }, 0);
    return new Money(total, currency);
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

let five = new Money(5, "USD")
let ten = new Money(10, "USD")
let fifteen = new Money(15, "USD")
let portfolio = new Portfolio();
portfolio.add(five, ten);
assert.deepStrictEqual(portfolio.evaluate("USD"), fifteen)
