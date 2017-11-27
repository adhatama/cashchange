package main

import (
	"fmt"
	"log"
)

type paymentChance struct {
	value  int
	detail []int
}

type cashChange struct {
	value  int
	detail []int
}

func main() {
	inputPrice := 0
	fmt.Print("Input price: ")
	_, err := fmt.Scanf("%d", &inputPrice)

	if err != nil {
		log.Fatal(err)
	}

	moneyBills := getIdr()
	moneyBills = reverse(moneyBills)

	fmt.Printf("Available money bills: %v\n", moneyBills)

	paymentChances := calculatePaymentChances(inputPrice, moneyBills)

	fmt.Println("=== Payment Chances ===")
	for _, val := range paymentChances {
		fmt.Printf("%d => %v\n", val.value, val.detail)
	}

	cashChanges := calculateCashChanges(inputPrice, paymentChances, moneyBills)

	fmt.Println("=== Cash Change Possibilites ===")
	for _, val := range cashChanges {
		fmt.Printf("%d => %v\n", val.value, val.detail)
	}
}

func getIdr() (x []int) {
	moneyBills := []int{500, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

	return moneyBills
}

func isPaymentChancesFound(result int, paymentChances []paymentChance) (found bool) {
	for _, val := range paymentChances {
		if val.value == result {
			return true
		}
	}

	return false
}

func sum(val []int) (result int) {
	result = 0
	for _, v := range val {
		result += v
	}
	return
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func calculatePaymentChances(inputPrice int, moneyBills []int) (paymentChances []paymentChance) {
	paymentChances = []paymentChance{}
	tempPChance := paymentChance{}

	for i := 0; i < len(moneyBills)-1; {
		sumDetail := sum(tempPChance.detail)

		if sumDetail == inputPrice {
			break
		}

		if inputPrice > sumDetail {
			tempPChance.detail = append(tempPChance.detail, moneyBills[i])
		} else if isPaymentChancesFound(sumDetail, paymentChances) {
			tempPChance.detail = tempPChance.detail[:len(tempPChance.detail)-1]
			i++
		} else {
			tempPChance.value = sumDetail
			paymentChances = append(paymentChances, tempPChance)
			tempPChance = paymentChance{}
			i = 0
		}
	}

	return
}

func calculateCashChanges(inputPrice int, paymentChances []paymentChance, moneyBills []int) (cashChanges []cashChange) {
	cashChanges = []cashChange{}
	tempCashChange := cashChange{}

	for _, pChance := range paymentChances {
		remaining := pChance.value - inputPrice

		for i := 0; i < len(moneyBills)-1; {
			if remaining == 0 {
				break
			}

			if remaining >= moneyBills[i] {
				tempCashChange.detail = append(tempCashChange.detail, moneyBills[i])
				remaining -= moneyBills[i]
			} else {
				i++
			}
		}

		tempCashChange.value = pChance.value
		cashChanges = append(cashChanges, tempCashChange)
		tempCashChange = cashChange{}
	}

	return
}
