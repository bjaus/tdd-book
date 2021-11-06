from __future__ import annotations

import typing as t
import unittest


class Money:
    def __init__(self, amount: int | float, currency: str):
        self.amount = amount
        self.currency = currency

    def __eq__(self, other: Money):
        return self.amount == other.amount and self.currency == other.currency

    def times(self, multiplier: int | float):
        return Money(amount=self.amount * multiplier, currency=self.currency)

    def divide(self, divisor: int | float):
        return Money(amount=self.amount / divisor, currency=self.currency)


class TestMoney(unittest.TestCase):
    def test_multiplication_in_dollars(self):
        money = Money(amount=5, currency="USD")
        expected = Money(amount=10, currency="USD")
        self.assertEqual(money.times(2), expected)

    def test_multiplication_in_euros(self):
        money = Money(amount=10, currency="EUR")
        expected = Money(amount=20, currency="EUR")
        self.assertEqual(money.times(2), expected)

    def test_division(self):
        money = Money(amount=4002, currency="KRW")
        expected = Money(amount=1000.5, currency="KRW")
        self.assertEqual(money.divide(4), expected)


if __name__ == "__main__":
    unittest.main()
