package main

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
	}

	input3 := 84000
	testData3 := map[int][]int{
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

}

func isValidPaymentChances(paymentChances []paymentChance, testData map[int][]int) (valid bool) {
	for key, val := range testData {
		valid = false

		for _, paymentChance := range paymentChances {
			if paymentChance.value == key {
				valid = reflect.DeepEqual(paymentChance.detail, val)
				break
			}
		}
	}

	return
}
