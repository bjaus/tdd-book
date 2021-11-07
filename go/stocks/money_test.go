package stocks_test

import (
	"reflect"
	s "tdd/stocks"
	"testing"
)

var bank s.Bank

func init() {
	bank = s.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	bank := s.NewBank()
	tenEuros := s.NewMoney(10, "EUR")
	result, err := bank.Convert(tenEuros, "Kalganid")
	assertNil(t, result)
	assertEqual(t, "EUR->Kalganid", err.Error())
}

func TestConversion(t *testing.T) {
	bank := s.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	tenEuros := s.NewMoney(10, "EUR")
	result, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, *result, s.NewMoney(12, "USD"))
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	var portfolio s.Portfolio

	oneDollar := s.NewMoney(1, "USD")
	oneEuro := s.NewMoney(1, "EUR")
	oneWon := s.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar, oneEuro, oneWon)

	expected := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	value, result := portfolio.Evaluate(bank, "Kalganid")

	assertNil(t, value)
	assertEqual(t, expected, result.Error())
}

func TestAddtiionaOfDollarsAndWons(t *testing.T) {
	var portfolio s.Portfolio

	one := s.NewMoney(1, "USD")
	eleven := s.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(one, eleven)
	result, _ := portfolio.Evaluate(bank, "KRW")

	expected := s.NewMoney(2200, "KRW")

	assertEqual(t, expected, *result)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio

	five := s.NewMoney(5, "USD")
	ten := s.NewMoney(10, "EUR")

	portfolio = portfolio.Add(five, ten)

	expected := s.NewMoney(17, "USD")
	result, _ := portfolio.Evaluate(bank, "USD")

	assertEqual(t, expected, *result)
}

func TestAddition(t *testing.T) {
	var portfolio s.Portfolio

	five := s.NewMoney(5, "USD")
	ten := s.NewMoney(10, "USD")
	expected := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(five, ten)
	result, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, expected, *result)
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

func assertNil(t *testing.T, value interface{}) {
	if value != nil && !reflect.ValueOf(value).IsNil() {
		t.Errorf("Expected error to be nil, got %+v", value)
	}
}
