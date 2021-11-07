from money import Money
from bank import Bank


class Portfolio:
    def __init__(self):
        self.__moneys = []

    def add(self, *moneys: tuple[Money]):
        self.__moneys.extend(moneys)

    def evaluate(self, bank: Bank, currency: str):
        total = 0.0
        failures = []
        for m in self.__moneys:
            try:
                total += bank.convert(m, currency).amount
            except Exception as e:
                failures.append(e)

        if not failures:
            return Money(total, currency)

        failure_message = ",".join(i.args[0] for i in failures)
        raise Exception("Missing exchange rate(s):[" + failure_message + "]")
