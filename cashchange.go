package cashchange

// CurrentCurrency config
var CurrentCurrency = "IDR"

// PaymentChance is to define structure data of payment chance possibilities
type PaymentChance struct {
	Value  int
	Detail []int
}

// CashChange is to define structure data of cash change possibilities
type CashChange struct {
	Value  int
	Detail []int
}

// Get is the main function of this package.
// Use this method to get the possibilities of cash change.
func Get(inputPrice int) []CashChange {
	if inputPrice <= 0 {
		return []CashChange{}
	}

	rounding(&inputPrice)

	moneyBills := GetCurrencyBills(CurrentCurrency)

	paymentChances := calculatePaymentChances(inputPrice, moneyBills)

	return calculateCashChanges(inputPrice, paymentChances, moneyBills)
}

// GetPaymentChances is to find the customer payment possibilities with details
// without the cash change possibilities
func GetPaymentChances(inputPrice int) []PaymentChance {
	if inputPrice <= 0 {
		return []PaymentChance{}
	}

	rounding(&inputPrice)

	moneyBills := GetCurrencyBills(CurrentCurrency)

	return calculatePaymentChances(inputPrice, moneyBills)
}

// GetCurrencyBills is used to get slice of currency bills based on params.
// Still only supports IDR
func GetCurrencyBills(currency string) []int {
	if currency == "" {
		currency = "IDR"
	}

	switch currency {
	case "IDR":
		return getIdr()
	default:
		return []int{}
	}
}

func getIdr() (x []int) {
	moneyBills := []int{100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

	return reverse(moneyBills)
}

func isPaymentChancesFound(result int, paymentChances []PaymentChance) (found bool) {
	for _, val := range paymentChances {
		if val.Value == result {
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

func rounding(val *int) {
	divider := 100
	mod := *val % divider

	if mod != 0 {
		*val = *val + divider - mod
	}
}

func calculatePaymentChances(inputPrice int, moneyBills []int) (paymentChances []PaymentChance) {
	paymentChances = []PaymentChance{}
	tempPChance := PaymentChance{}

	for i := 0; i < len(moneyBills)-1; {
		sumDetail := sum(tempPChance.Detail)

		if sumDetail == inputPrice {
			break
		}

		if inputPrice > sumDetail {
			tempPChance.Detail = append(tempPChance.Detail, moneyBills[i])
		} else if isPaymentChancesFound(sumDetail, paymentChances) {
			tempPChance.Detail = tempPChance.Detail[:len(tempPChance.Detail)-1]
			i++
		} else {
			tempPChance.Value = sumDetail
			paymentChances = append(paymentChances, tempPChance)
			tempPChance = PaymentChance{}
			i = 0
		}
	}

	return
}

func calculateCashChanges(inputPrice int, paymentChances []PaymentChance, moneyBills []int) (cashChanges []CashChange) {
	cashChanges = []CashChange{}
	tempCashChange := CashChange{}

	for _, pChance := range paymentChances {
		remaining := pChance.Value - inputPrice

		for i := 0; i < len(moneyBills); {
			if remaining == 0 {
				break
			}

			if remaining >= moneyBills[i] {
				tempCashChange.Detail = append(tempCashChange.Detail, moneyBills[i])
				remaining -= moneyBills[i]
			} else {
				i++
			}
		}

		tempCashChange.Value = pChance.Value
		cashChanges = append(cashChanges, tempCashChange)
		tempCashChange = CashChange{}
	}

	return
}
