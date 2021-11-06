package stocks_test

import (
	s "tdd/stocks"
	"testing"
)

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio

	five := s.NewMoney(5, "USD")
	ten := s.NewMoney(10, "EUR")

	portfolio = portfolio.Add(five, ten)

	expected := s.NewMoney(17, "USD")
	result := portfolio.Evaluate("USD")

	assertEqual(t, expected, result)
}

func TestAddition(t *testing.T) {
	var portfolio s.Portfolio

	five := s.NewMoney(5, "USD")
	ten := s.NewMoney(10, "USD")
	expected := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(five, ten)
	result := portfolio.Evaluate("USD")

	assertEqual(t, expected, result)
}

func TestDivision(t *testing.T) {
	money := s.NewMoney(4002, "KRW")
	result := money.Divide(4)
	expected := s.NewMoney(1000.5, "KRW")
	assertEqual(t, expected, result)
}

func TestMuliplication(t *testing.T) {
	money := s.NewMoney(10, "EUR")
	result := money.Times(2)
	expected := s.NewMoney(20, "EUR")
	assertEqual(t, expected, result)
}

func assertEqual(t *testing.T, expected s.Money, actual s.Money) {
	if expected != actual {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
