package stocks

type Portfolio []Money

func (p Portfolio) Add(moneys ...Money) Portfolio {
	for _, m := range moneys {
		p = append(p, m)
	}
	return p
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total += convert(m, currency)
	}
	return NewMoney(total, currency)
}

func convert(money Money, currency string) float64 {
	key := money.currency + "->" + currency
	rate := exchangeRates[key]
	if money.currency == currency {
		return money.amount
	}
	return money.amount * rate
}
