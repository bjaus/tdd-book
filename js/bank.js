const Money = require("./money");

class Bank {
  constructor() {
    this.exchangeRates = new Map();
  }

  makeKey(from, to) {
    return from + "->" + to
  }

  addExchangeRate(from, to, rate) {
    let key = this.makeKey(from, to);
    this.exchangeRates.set(key, rate);
  }

  convert(money, currency) {
    if (money.currency === currency) {
      return new Money(money.amount, money.currency);
    }
    let key = this.makeKey(money.currency, currency);
    let rate = this.exchangeRates.get(key);
    if (rate === undefined) {
      throw new Error(key);
    }
    return new Money(money.amount * rate, currency);
  }
}

module.exports = Bank;
