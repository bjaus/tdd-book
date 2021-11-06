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
		total += m.amount
	}
	return NewMoney(total, currency)
}
