package main

import (
	"fmt"
	"math"
)

func CalculateMonthlyRepaymentAmount() {
	loanAmount := Ask("Enter loan amount", 1000.0)
	interestRate := Ask("Enter interest rate", 0.0)
	months := Ask("Enter months", 0.2)
	depositAmount := Ask("Enter deposit amount", 0.0)

	if depositAmount >= loanAmount {
		fmt.Println("Deposit amount is greater than or same as loan amount.")
		fmt.Println("--")
		CalculateMonthlyRepaymentAmount()
	}

	amountToPay := loanAmount - depositAmount	
	monthlyInterestRate := interestRate/100/12
	
	// @see https://www.investopedia.com/terms/a/amortization.asp
	monthlyAmount := amountToPay * ((monthlyInterestRate * math.Pow(1 + monthlyInterestRate, months))/(math.Pow(1 + monthlyInterestRate, months) - 1))

	output := fmt.Sprintf("Monthly repayment amount is %.2f", monthlyAmount)

	DisplayInfo(output)
}