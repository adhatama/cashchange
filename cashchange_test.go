package cashchange

import (
	"reflect"
	"testing"
)

func TestCurrencyBills(t *testing.T) {
	// Given
	currency := "IDR"
	testData1 := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500}

	// When
	currencyBills1 := GetCurrencyBills(currency)
	currencyBills2 := GetCurrencyBills("")

	// Then
	if !reflect.DeepEqual(currencyBills1, testData1) {
		t.Error("Expected: ", testData1, "Given: ", currencyBills1)
	}
	if !reflect.DeepEqual(currencyBills2, testData1) {
		t.Error("Expected: ", testData1, "Given: ", currencyBills2)
	}
}

func TestCalculatePaymentChances(t *testing.T) {
	// Given
	input1 := 76000
	testData1 := map[int][]int{
		77000:  []int{50000, 20000, 5000, 2000},
		80000:  []int{50000, 20000, 10000},
		90000:  []int{50000, 20000, 20000},
		100000: []int{100000},
	}

	input2 := 140000
	testData2 := map[int][]int{
		150000: []int{100000, 50000},
		200000: []int{100000, 100000},
	}

	input3 := 83000
	testData3 := map[int][]int{
		84000:  []int{50000, 20000, 10000, 2000, 2000},
		85000:  []int{50000, 20000, 10000, 5000},
		90000:  []int{50000, 20000, 20000},
		100000: []int{100000},
	}

	input4 := 0
	testData4 := map[int][]int{}

	input5 := -90000
	testData5 := map[int][]int{}

	// When
	paymentChances1 := GetPaymentChances(input1)
	paymentChances2 := GetPaymentChances(input2)
	paymentChances3 := GetPaymentChances(input3)
	paymentChances4 := GetPaymentChances(input4)
	paymentChances5 := GetPaymentChances(input5)

	// Then
	if !isValidPaymentChances(paymentChances1, testData1) {
		t.Error("Expected: ", testData1, "Given: ", paymentChances1)
	}
	if !isValidPaymentChances(paymentChances2, testData2) {
		t.Error("Expected: ", testData2, "Given: ", paymentChances2)
	}
	if !isValidPaymentChances(paymentChances3, testData3) {
		t.Error("Expected: ", testData3, "Given: ", paymentChances3)
	}
	if !isValidPaymentChances(paymentChances4, testData4) {
		t.Error("Expected: ", testData4, "Given: ", paymentChances4)
	}
	if !isValidPaymentChances(paymentChances5, testData5) {
		t.Error("Expected: ", testData5, "Given: ", paymentChances5)
	}
}

func TestCalculateCashChanges(t *testing.T) {
	// Given
	input1 := 76000
	testData1 := map[int][]int{
		77000:  []int{1000},
		80000:  []int{2000, 2000},
		90000:  []int{10000, 2000, 2000},
		100000: []int{20000, 2000, 2000},
	}

	input2 := 140000
	testData2 := map[int][]int{
		150000: []int{10000},
		200000: []int{50000, 10000},
	}

	input3 := 83000
	testData3 := map[int][]int{
		84000:  []int{1000},
		85000:  []int{2000},
		90000:  []int{5000, 2000},
		100000: []int{10000, 5000, 2000},
	}

	input4 := 0
	testData4 := map[int][]int{}

	input5 := -90000
	testData5 := map[int][]int{}

	// When
	cashChanges1 := Get(input1)
	cashChanges2 := Get(input2)
	cashChanges3 := Get(input3)
	cashChanges4 := Get(input4)
	cashChanges5 := Get(input5)

	// Then
	if !isValidCashChanges(cashChanges1, testData1) {
		t.Error("Expected: ", testData1, "Given: ", cashChanges1)
	}
	if !isValidCashChanges(cashChanges2, testData2) {
		t.Error("Expected: ", testData2, "Given: ", cashChanges2)
	}
	if !isValidCashChanges(cashChanges3, testData3) {
		t.Error("Expected: ", testData3, "Given: ", cashChanges3)
	}
	if !isValidCashChanges(cashChanges4, testData4) {
		t.Error("Expected: ", testData4, "Given: ", cashChanges4)
	}
	if !isValidCashChanges(cashChanges5, testData5) {
		t.Error("Expected: ", testData5, "Given: ", cashChanges5)
	}
}

func isValidPaymentChances(paymentChances []PaymentChance, testData map[int][]int) (valid bool) {
	for key, val := range testData {
		valid = false

		for _, paymentChance := range paymentChances {
			if paymentChance.Value == key {
				valid = reflect.DeepEqual(paymentChance.Detail, val)
				break
			}
		}

		if !valid {
			return false
		}
	}

	return true
}

func isValidCashChanges(cashChanges []CashChange, testData map[int][]int) (valid bool) {
	for key, val := range testData {
		valid = false

		for _, cashChange := range cashChanges {
			if cashChange.Value == key {
				valid = reflect.DeepEqual(cashChange.Detail, val)
				break
			}
		}

		if !valid {
			return false
		}
	}

	return true
}
