import unittest

from money import Money
from portfolio import Portfolio
from bank import Bank


class TestMoney(unittest.TestCase):
    def setUp(self):
        self.bank = Bank()
        self.bank.add_exchange_rate("EUR", "USD", 1.2)
        self.bank.add_exchange_rate("USD", "KRW", 1100)

    def test_conversion_with_missing_exchange_rate(self):
        bank = Bank()
        ten_euros = Money(10, "EUR")
        with self.assertRaisesRegex(Exception, "EUR->Kalganid"):
            bank.convert(ten_euros, "Kalganid")

    def test_conversion(self):
        bank = Bank()
        bank.add_exchange_rate("EUR", "USD", 1.2)
        ten_euros = Money(10, "EUR")
        self.assertEqual(bank.convert(ten_euros, "USD"), Money(12, "USD"))

    def test_addition_with_multiple_missing_exchange_rates(self):
        one_dollar = Money(1, "USD")
        one_euro = Money(1, "EUR")
        one_won = Money(1, "KRW")
        portfolio = Portfolio()
        portfolio.add(one_dollar, one_euro, one_won)

        with self.assertRaisesRegex(
            Exception,
            r"Missing exchange rate\(s\):\[USD\-\>Kalganid,EUR\-\>Kalganid,KRW\-\>Kalganid\]",
        ):
            portfolio.evaluate(Bank(), "Kalganid")

    def test_addition_of_dollars_and_wons(self):
        one = Money(1, "USD")
        eleven = Money(1100, "KRW")
        portfolio = Portfolio()
        portfolio.add(one, eleven)
        expected = Money(2200, "KRW")
        self.assertEqual(expected, portfolio.evaluate(self.bank, "KRW"))

    def test_addition_of_dollars_and_euros(self):
        five = Money(5, "USD")
        ten = Money(10, "EUR")
        portfolio = Portfolio()
        portfolio.add(five, ten)
        expected = Money(17, "USD")
        result = portfolio.evaluate(self.bank, "USD")
        self.assertEqual(expected, result)

    def test_addition(self):
        five = Money(amount=5, currency="USD")
        ten = Money(amount=10, currency="USD")
        fifteen = Money(amount=15, currency="USD")
        portfolio = Portfolio()
        portfolio.add(five, ten)
        self.assertEqual(fifteen, portfolio.evaluate(self.bank, "USD"))

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
