package main

import (
	"testing"
)

func assertMoneyIsEqual(t *testing.T, expected Money, actual Money) {
	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestMultiplicationinDollars(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}
	actualTenner := fiver.Times(2)
	expectedTenner := Money{
		amount:   10,
		currency: "USD",
	}
	assertMoneyIsEqual(t, expectedTenner, actualTenner)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualTwentyEuros := tenEuros.Times(2)
	expectedTwentyEuros := Money{amount: 20, currency: "EUR"}
	assertMoneyIsEqual(t, expectedTwentyEuros, actualTwentyEuros)
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	assertMoneyIsEqual(t, expectedMoneyAfterDivision, actualMoneyAfterDivision)
}

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}
