from money import Money


class Bank:
    def __init__(self):
        self.__rates = {}

    @staticmethod
    def __make_key(from_, to):
        return f"{from_}->{to}"

    def add_exchange_rate(self, from_, to, rate):
        key = self.__make_key(from_, to)
        self.__rates[key] = rate

    def convert(self, money, currency):
        if money.currency == currency:
            return Money(money.amount, currency)

        key = self.__make_key(money.currency, currency)
        try:
            return Money(money.amount * self.__rates[key], currency)
        except KeyError:
            raise Exception(key)
