const Money = require("./money");
const exchangeRates = require("./rates");

class Portfolio {
  constructor() {
    this.moneys = [];
  }

  add(...moneys) {
    this.moneys = this.moneys.concat(moneys);
  }

  evaluate(currency) {
    let total = this.moneys.reduce((sum, money) => {
      return sum + this.convert(money, currency);
    }, 0);
    return new Money(total, currency);
  }

  convert(money, currency) {
    let key = money.currency + '->' + currency;
    let rate = exchangeRates.get(key);
    if (money.currency === currency) {
      return money.amount;
    }
    return money.amount * rate;
  }
}

module.exports = Portfolio;
