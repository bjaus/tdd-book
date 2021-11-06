import functools
import operator

from money import Money


class Portfolio:
    def __init__(self):
        self.moneys = []

    def add(self, *moneys: tuple[Money]):
        self.moneys.extend(moneys)

    def evaluate(self, currency: str):
        total = functools.reduce(operator.add, map(lambda m: m.amount, self.moneys))
        return Money(amount=total, currency=currency)
