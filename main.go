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
	_, err := fmt.Scanf("%d", &inputPrice)

	if err != nil {
		log.Fatal(err)
	}

	moneyBills := getIdr()

	for i, j := 0, len(moneyBills)-1; i < j; i, j = i+1, j-1 {
		moneyBills[i], moneyBills[j] = moneyBills[j], moneyBills[i]
	}

	fmt.Println(moneyBills)

	paymentChances := []paymentChance{}
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

	fmt.Println("=== Payment Chances ===")
	for _, val := range paymentChances {
		fmt.Printf("%d => %v\n", val.value, val.detail)
	}

	cashChanges := []paymentChance{}
	tempPChance = paymentChance{}
	for _, pChance := range paymentChances {
		remaining := pChance.value - inputPrice

		for i := 0; i < len(moneyBills)-1; {
			if remaining == 0 {
				break
			}

			if remaining >= moneyBills[i] {
				tempPChance.detail = append(tempPChance.detail, moneyBills[i])
				remaining -= moneyBills[i]
			} else {
				i++
			}
		}

		tempPChance.value = pChance.value
		cashChanges = append(cashChanges, tempPChance)
		tempPChance = paymentChance{}
	}

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
