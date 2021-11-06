package main

import "testing"

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
