package stocks

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) Money {
	return Money{amount: amount, currency: currency}
}

func (m Money) Times(multiplier float64) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) Divide(divisor float64) Money {
	return Money{amount: m.amount / divisor, currency: m.currency}
}
