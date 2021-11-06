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

module.exports = Money;
