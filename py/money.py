from __future__ import annotations


class Money:
    def __init__(self, amount: int | float, currency: str):
        self.amount = amount
        self.currency = currency

    def __repr__(self):
        return f"{self.amount:0.2f} {self.currency}"

    def __eq__(self, other: Money):
        return self.amount == other.amount and self.currency == other.currency

    def times(self, multiplier: int | float):
        return Money(amount=self.amount * multiplier, currency=self.currency)

    def divide(self, divisor: int | float):
        return Money(amount=self.amount / divisor, currency=self.currency)
