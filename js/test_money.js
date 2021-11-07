const assert = require('assert');
const Money = require('./money');
const Portfolio = require('./portfolio');
const Bank = require('./bank');


class MoneyTest {
  testConversionWithMissingExchangeRate() {
    let bank = new Bank();
    let tenEuros = new Money(10, "EUR");
    let expected = new Error("EUR->Kalganid");
    assert.throws(function () { bank.convert(tenEuros, "Kalganid"); }, expected);
  }

  testConversionWithDifferentRatesBetweenTwoCurrencies() {
    let tenEuros = new Money(10, "EUR");
    assert.deepStrictEqual(this.bank.convert(tenEuros, "USD"), new Money(12, "USD"));

    this.bank.addExchangeRate("EUR", "USD", 1.3);
    assert.deepStrictEqual(this.bank.convert(tenEuros, "USD"), new Money(13, "USD"));
  }

  testAdditionWithMultipleMissingExchangeRates() {
    let oneDollar = new Money(1, "USD")
    let oneEuro = new Money(1, "EUR")
    let oneWon = new Money(1, "KRW")
    let portfolio = new Portfolio();
    portfolio.add(oneDollar, oneEuro, oneWon);
    let expected = new Error("Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid]");
    assert.throws(function() {portfolio.evaluate(new Bank(), "Kalganid")}, expected);
  }

  testAdditionOfDollarsAndWons() {
    let one = new Money(1, "USD")
    let eleven = new Money(1100, "KRW")
    let portfolio = new Portfolio();
    portfolio.add(one, eleven);
    let expected = new Money(2200, "KRW")
    assert.deepStrictEqual(portfolio.evaluate(this.bank, "KRW"), expected);
  }

  testAdditionOfDollarsAndEuros() {
    let five = new Money(5, "USD")
    let ten = new Money(10, "EUR")
    let portfolio = new Portfolio();
    portfolio.add(five, ten)
    let expected = new Money(17, "USD")
    assert.deepStrictEqual(portfolio.evaluate(this.bank, "USD"), expected)
  }

  testMultiplication() {
    let money = new Money(10, "EUR")
    let expected = new Money(20, "EUR")
    assert.deepStrictEqual(money.times(2), expected);
  }

  testDivision() {
    let money = new Money(4002, "KRW")
    let expected = new Money(1000.5, "KRW")
    assert.deepStrictEqual(money.divide(4), expected)
  }

  testAddition() {
    let five = new Money(5, "USD")
    let ten = new Money(10, "USD")
    let fifteen = new Money(15, "USD")
    let portfolio = new Portfolio();
    portfolio.add(five, ten);
    assert.deepStrictEqual(portfolio.evaluate(this.bank, "USD"), fifteen)
  }

  randomize(methods) {
    for (let i = methods.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [methods[i], methods[j]] = [methods[j], methods[i]];
    }
    return methods;
  }

  get() {
    let proto = MoneyTest.prototype;
    let props = Object.getOwnPropertyNames(proto);
    let methods = props.filter(p => {
      return typeof proto[p] === "function" && p.startsWith("test");
    });
    return this.randomize(methods);
  }

  setUp() {
    this.bank = new Bank();
    this.bank.addExchangeRate("EUR", "USD", 1.2);
    this.bank.addExchangeRate("USD", "KRW", 1100);
  }

  run() {
    let testMethods = this.get();
    testMethods.forEach(m => {
      console.log("Running: %s()", m)
      let method = Reflect.get(this, m);
      try {
        this.setUp();
        Reflect.apply(method, this, []);
      } catch (e) {
        if (e instanceof assert.AssertionError) {
          console.log(e);
        } else {
          throw e;
        }
      }
    })
  }
}

new MoneyTest().run();
