const assert = require('assert');
const Money = require('./money');
const Portfolio = require('./portfolio');


class MoneyTest {
  testAdditionOfDollarsAndWons() {
    let one = new Money(1, "USD")
    let eleven = new Money(1100, "KRW")
    let portfolio = new Portfolio();
    portfolio.add(one, eleven);
    let expected = new Money(2200, "KRW")
    assert.deepStrictEqual(portfolio.evaluate("KRW"), expected);
  }

  testAdditionOfDollarsAndEuros() {
    let five = new Money(5, "USD")
    let ten = new Money(10, "EUR")
    let portfolio = new Portfolio();
    portfolio.add(five, ten)
    let expected = new Money(17, "USD")
    assert.deepStrictEqual(portfolio.evaluate("USD"), expected)
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
    assert.deepStrictEqual(portfolio.evaluate("USD"), fifteen)
  }

  get() {
    let proto = MoneyTest.prototype;
    let props = Object.getOwnPropertyNames(proto);
    let methods = props.filter(p => {
      return typeof proto[p] === "function" && p.startsWith("test");
    });
    return methods;
  }

  run() {
    let testMethods = this.get();
    testMethods.forEach(m => {
      console.log("Running: %s()", m)
      let method = Reflect.get(this, m);
      try {
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
