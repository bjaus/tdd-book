from money import Money
from rates import EXCHANGE_RATES


class Portfolio:
    def __init__(self):
        self.__moneys = []

    def add(self, *moneys: tuple[Money]):
        self.__moneys.extend(moneys)

    def evaluate(self, currency: str):
        total = sum(self.__convert(i, currency) for i in self.__moneys)
        return Money(amount=total, currency=currency)

    def __convert(self, money: Money, currency: str) -> int | float:
        key = f"{money.currency}->{currency}"
        rate = EXCHANGE_RATES.get(key)
        if money.currency == currency:
            return money.amount
        return money.amount * rate
