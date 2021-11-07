package stocks_test

import (
	s "tdd/stocks"
	"testing"
)

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	oneEuro := s.NewMoney(1, "EUR")
	oneWon := s.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar, oneEuro, oneWon)

	expected := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	_, result := portfolio.Evaluate("Kalganid")

	assertEqual(t, expected, result.Error())
}

func TestAddtiionaOfDollarsAndWons(t *testing.T) {
	var portfolio s.Portfolio

	one := s.NewMoney(1, "USD")
	eleven := s.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(one, eleven)
	result, _ := portfolio.Evaluate("KRW")

	expected := s.NewMoney(2200, "KRW")

	assertEqual(t, expected, result)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio

	five := s.NewMoney(5, "USD")
	ten := s.NewMoney(10, "EUR")

	portfolio = portfolio.Add(five, ten)

	expected := s.NewMoney(17, "USD")
	result, _ := portfolio.Evaluate("USD")

	assertEqual(t, expected, result)
}

func TestAddition(t *testing.T) {
	var portfolio s.Portfolio

	five := s.NewMoney(5, "USD")
	ten := s.NewMoney(10, "USD")
	expected := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(five, ten)
	result, _ := portfolio.Evaluate("USD")

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

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
