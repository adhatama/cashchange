package main

import (
	"fmt"
	"log"

	"github.com/adhatama/cashchange"
)

func main() {
	input := 0
	fmt.Print("Input price: ")
	_, err := fmt.Scanf("%d", &input)

	if err != nil {
		log.Fatal(err)
	}

	currencyBills := cashchange.GetCurrencyBills("IDR")
	fmt.Printf("Available currency bills: %v\n", currencyBills)

	paymentChances := cashchange.GetPaymentChances(input)

	fmt.Println("=== Payment Chances ===")
	for _, val := range paymentChances {
		fmt.Printf("%d => %v\n", val.Value, val.Detail)
	}

	cashChanges := cashchange.Get(input)

	fmt.Println("=== Cash Change Possibilites ===")
	for _, val := range cashChanges {
		fmt.Printf("%d => %v\n", val.Value, val.Detail)
	}
}
