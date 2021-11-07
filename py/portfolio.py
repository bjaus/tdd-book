from money import Money
from rates import EXCHANGE_RATES


class Portfolio:
    def __init__(self):
        self.__moneys = []

    def add(self, *moneys: tuple[Money]):
        self.__moneys.extend(moneys)

    def evaluate(self, currency: str):
        total = 0.0
        failures = []
        for m in self.__moneys:
            try:
                total += self.__convert(m, currency)
            except KeyError as e:
                failures.append(e)
        if not failures:
            return Money(total, currency)
        failure_message = ",".join(i.args[0] for i in failures)
        raise Exception("Missing exchange rate(s):[" + failure_message + "]")

    def __convert(self, money: Money, currency: str) -> int | float:
        if money.currency == currency:
            return money.amount
        key = f"{money.currency}->{currency}"
        rate = EXCHANGE_RATES[key]
        return money.amount * rate
