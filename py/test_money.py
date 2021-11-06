import unittest

from money import Money
from portfolio import Portfolio


class TestMoney(unittest.TestCase):
    def test_addition_of_dollars_and_euros(self):
        five = Money(5, "USD")
        ten = Money(10, "EUR")
        portfolio = Portfolio()
        portfolio.add(five, ten)
        expected = Money(17, "USD")
        result = portfolio.evaluate("USD")
        self.assertEqual(expected, result)

    def test_addition(self):
        five = Money(amount=5, currency="USD")
        ten = Money(amount=10, currency="USD")
        fifteen = Money(amount=15, currency="USD")
        portfolio = Portfolio()
        portfolio.add(five, ten)
        self.assertEqual(fifteen, portfolio.evaluate("USD"))

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
