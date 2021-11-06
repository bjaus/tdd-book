package main

import "testing"

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
	return Money{amount: total, currency: currency}
}

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier float64) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}

func (m Money) Divide(divisor float64) Money {
	return Money{amount: m.amount / divisor, currency: m.currency}
}

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioUSD Money

	five := Money{amount: 5, currency: "USD"}
	ten := Money{amount: 10, currency: "USD"}
	fifteen := Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(five, ten)
	portfolioUSD = portfolio.Evaluate("USD")

	assertEqual(t, fifteen, portfolioUSD)
}

func TestDivision(t *testing.T) {
	money := Money{amount: 4002, currency: "KRW"}
	actual := money.Divide(4)
	expected := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expected, actual)
}

func TestMuliplicationInDollars(t *testing.T) {
	money := Money{amount: 5, currency: "USD"}
	actual := money.Times(2)
	expected := Money{amount: 10, currency: "USD"}
	assertEqual(t, expected, actual)
}

func TestMuliplicationInEuros(t *testing.T) {
	money := Money{amount: 10, currency: "EUR"}
	actual := money.Times(2)
	expected := Money{amount: 20, currency: "EUR"}
	assertEqual(t, expected, actual)
}
