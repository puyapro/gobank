package main

import (
	"fmt"
	"gobank/loan"
)

func main() {
	bankRules := loan.Rules{
		MinAge:    18,
		MinIncome: 1000,
	}

	if loan.Check(age, income, hasBadCredit, bankRules) {
		fmt.Println("loan approved")
	} else {
		fmt.Println("loan rejected")
	}
}
