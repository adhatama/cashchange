package main

import (
	"fmt"
	"log"
)

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

	paymentChances := make(map[int][]int)
	var tempResult []int
	for i := 0; i < len(moneyBills)-1; {
		sumTempResult := sum(tempResult)

		if sumTempResult == inputPrice {
			break
		}

		if inputPrice > sumTempResult {
			tempResult = append(tempResult, moneyBills[i])
		} else if isPaymentChancesFound(sumTempResult, paymentChances) {
			tempResult = tempResult[:len(tempResult)-1]
			i++
		} else {
			paymentChances[sumTempResult] = tempResult
			tempResult = nil
			i = 0
		}
	}

	fmt.Println("=== Payment Chances ===")
	for key, val := range paymentChances {
		fmt.Printf("%d => ", key)
		fmt.Println(val)
	}

	results := make(map[int][]int)
	tempResult = nil
	for key := range paymentChances {
		remaining := key - inputPrice

		for i := 0; i < len(moneyBills)-1; {
			if remaining == 0 {
				break
			}

			if remaining >= moneyBills[i] {
				tempResult = append(tempResult, moneyBills[i])
				remaining -= moneyBills[i]
			} else {
				i++
			}
		}

		results[key] = tempResult
		tempResult = nil
	}

	fmt.Println("=== Cash Change Possibilites ===")
	for key, val := range results {
		fmt.Printf("%d => ", key)
		fmt.Println(val)
	}
}

func getIdr() (x []int) {
	moneyBills := []int{500, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

	return moneyBills
}

func isPaymentChancesFound(result int, paymentChances map[int][]int) (found bool) {
	for key := range paymentChances {
		if key == result {
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
