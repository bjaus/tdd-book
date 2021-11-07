package stocks

import "errors"

type Bank struct {
	exchangeRates map[string]float64
}

func makeKey(from, to string) string {
	return from + "->" + to
}

func (b Bank) AddExchangeRate(currencyFrom string, currencyTo string, rate float64) {
	key := makeKey(currencyFrom, currencyTo)
	b.exchangeRates[key] = rate
}

func (b Bank) Convert(money Money, currencyTo string) (*Money, error) {
	if money.currency == currencyTo {
		result := NewMoney(money.amount, money.currency)
		return &result, nil
	}
	key := makeKey(money.currency, currencyTo)
	rate, ok := b.exchangeRates[key]
	if ok {
		result := NewMoney(money.amount*rate, currencyTo)
		return &result, nil
	}
	return nil, errors.New(key)
}

func NewBank() Bank {
	return Bank{exchangeRates: make(map[string]float64)}
}
