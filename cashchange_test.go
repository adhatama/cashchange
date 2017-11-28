package cashchange

import (
	"reflect"
	"testing"
)

func TestCalculatePaymentChances(t *testing.T) {
	// Given
	moneyBills := getIdr()

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

	// When
	paymentChances1 := calculatePaymentChances(input1, moneyBills)
	paymentChances2 := calculatePaymentChances(input2, moneyBills)
	paymentChances3 := calculatePaymentChances(input3, moneyBills)

	// Then
	if !isValidPaymentChances(paymentChances1, testData1) {
		t.Error("Not working")
	}
	if !isValidPaymentChances(paymentChances2, testData2) {
		t.Error("Not working")
	}
	if !isValidPaymentChances(paymentChances3, testData3) {
		t.Error("Not working")
	}
}

func TestCalculateCashChanges(t *testing.T) {
	// Given
	moneyBills := getIdr()

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

	paymentChances1 := calculatePaymentChances(input1, moneyBills)
	paymentChances2 := calculatePaymentChances(input2, moneyBills)
	paymentChances3 := calculatePaymentChances(input3, moneyBills)

	// When
	cashChanges1 := calculateCashChanges(input1, paymentChances1, moneyBills)
	cashChanges2 := calculateCashChanges(input2, paymentChances2, moneyBills)
	cashChanges3 := calculateCashChanges(input3, paymentChances3, moneyBills)

	// Then
	if !isValidCashChanges(cashChanges1, testData1) {
		t.Error("Not working")
	}
	if !isValidCashChanges(cashChanges2, testData2) {
		t.Error("Not working")
	}
	if !isValidCashChanges(cashChanges3, testData3) {
		t.Error("Not working")
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
