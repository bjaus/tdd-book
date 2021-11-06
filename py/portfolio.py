from money import Money


class Portfolio:
    def __init__(self):
        self.__moneys = []
        self.__eur_to_usd = 1.2

    def add(self, *moneys: tuple[Money]):
        self.__moneys.extend(moneys)

    def evaluate(self, currency: str):
        total = sum(self.__convert(i, currency) for i in self.__moneys)
        return Money(amount=total, currency=currency)

    def __convert(self, money: Money, currency: str) -> int | float:
        if money.currency == currency:
            return money.amount
        return money.amount * self.__eur_to_usd
