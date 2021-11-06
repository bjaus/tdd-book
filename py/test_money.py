import typing as t
import unittest


class Dollar:
    def __init__(self, amount: int | float):
        self.amount = amount

    def times(self, multiplier: int | float):
        return Dollar(self.amount * multiplier)


class TestMoney(unittest.TestCase):
    def test_multiplication(self):
        fiver = Dollar(5)
        tenner = fiver.times(2)
        self.assertEqual(10, tenner.amount)


if __name__ == "__main__":
    unittest.main()
