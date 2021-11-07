import unittest

from money import Money
from portfolio import Portfolio


class TestMoney(unittest.TestCase):
    def test_addition_with_multiple_missing_exchange_rates(self):
        one_dollar = Money(1, "USD")
        one_euro = Money(1, "EUR")
        one_won = Money(1, "KRW")
        portfolio = Portfolio()
        portfolio.add(one_dollar, one_euro, one_won)

        with self.assertRaisesRegexp(
            Exception,
            r"Missing exchange rate\(s\):\[USD\-\>Kalganid,EUR\-\>Kalganid,KRW\-\>Kalganid\]",
        ):
            portfolio.evaluate("Kalganid")

    def test_addition_of_dollars_and_wons(self):
        one = Money(1, "USD")
        eleven = Money(1100, "KRW")
        portfolio = Portfolio()
        portfolio.add(one, eleven)
        expected = Money(2200, "KRW")
        self.assertEquals(expected, portfolio.evaluate("KRW"))

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
