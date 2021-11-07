const Money = require("./money");
const exchangeRates = require("./rates");

class Portfolio {
  constructor() {
    this.moneys = [];
  }

  add(...moneys) {
    this.moneys = this.moneys.concat(moneys);
  }

  evaluate(bank, currency) {
    let failures = [];
    let total = this.moneys.reduce((sum, money) => {
      try {
        let converted = bank.convert(money, currency);
        return sum + converted.amount;
      } catch (e) {
        failures.push(e.message);
        return sum;
      }
    }, 0);
    if (!failures.length) {
      return new Money(total, currency);
    }
    throw new Error("Missing exchange rate(s):[" + failures.join() + "]");
  }
}

module.exports = Portfolio;
