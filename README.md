# Cash Change
If you are somebody who sells something or a person who provide services, then you need to prepare your money so you can give your customer a proper change.  
Cash Change is here to help you prepare your money for change.

This apps is only available in IDR

## Installation
`go get github.com/adhatama/cashchange`

## Usage
- For the ready to use executable file, its in `cmd` folder
- The `main` function also already in `cmd` folder and showing how to use this package

```Go
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
```

## Case study
- You are pizza delivery man
    - The customer buy a pizza that cost IDR 140.000
    - You need to prepare money so you can give your customer a proper change
    - You input the cost of the pizza in this apps
    - The apps will tell you: 
        - if 150.000 then prepare 10.000
        - if 200.000 then prepare 50.000 & 10.000
        - So you need to bring a 10.000, a 20.000, and a 50.000 to prepare for every scenario your customer had
- You are gojek driver
    - The customer order gofood from you that cost IDR 76.000
    - You need to prepare money so you can give your customer a proper change
    - You input the cost to this apps
    - The apps will tell you:
        - if 77.000 then prepare 1.000
        - if 80.000 then prepare 2.000 & 2.000
        - if 90.000 then prepare 10.000 & 2.000 & 2.000
        - if 100.000 then prepare 20.000 & 2.000 & 2.000
        - So you need to bring a couple of 2.000, a 10.000, and a 20.000 to prepare for every scenario your customer had
- You are a cashier
    - You calculate the price from customer buying using your cashier apps
    - You got the total price, its IDR 83.000 and input that price to this apps
    - The apps will tell you:
        - If your customer pay 84.000 then give them change IDR 1.000
        - If 85.000 then give them change IDR 2.000
        - If 90.000 then give them change IDR 5.000 and 2.000
        - If 100.000 then give them change IDR 10.000, 5.000, and 2.000

## TODO
- [x] Handle changes up to the lowest possible currency bills, ex: IDR 100, 200
- [ ] Handle unchange input, ex: 96345