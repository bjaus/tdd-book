package stocks_test

import (
	"tdd/stocks"
	"testing"
)

func assertEqual(t *testing.T, expected stocks.Money, actual stocks.Money) {
	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestAddition(t *testing.T) {
	var portfolio stocks.Portfolio

	five := stocks.NewMoney(5, "USD")
	ten := stocks.NewMoney(10, "USD")
	expected := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(five, ten)
	result := portfolio.Evaluate("USD")

	assertEqual(t, expected, result)
}

func TestDivision(t *testing.T) {
	money := stocks.NewMoney(4002, "KRW")
	result := money.Divide(4)
	expected := stocks.NewMoney(1000.5, "KRW")
	assertEqual(t, expected, result)
}

func TestMuliplication(t *testing.T) {
	money := stocks.NewMoney(10, "EUR")
	result := money.Times(2)
	expected := stocks.NewMoney(20, "EUR")
	assertEqual(t, expected, result)
}
