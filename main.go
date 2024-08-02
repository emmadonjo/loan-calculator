package main

import (
	"fmt"
	"math"
)

func main() {
	run()	
}

func run() {
	selectAction := action()

    switch selectAction {
		case "1":
			calculateMonthlyRepaymentAmount()
		case "5":
			return
		default:
			run()
	}
}

func action() string {
	var selectAction string;

	fmt.Println("Select an action by entering the corresponding number. E.g. 1")
	fmt.Println("1 - Calculate monthly repayment amount")
	fmt.Println("2 - Calculate interest rate")
	fmt.Println("3 - Calculate number of months")
	fmt.Println("4 - Calculate loan amount")
	fmt.Println("5 - Exit")
	fmt.Print("Enter number: ")

	fmt.Scanln(&selectAction)

	if selectAction == "" {
		action()
	}

    return selectAction
}

func displayInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println()
	run()
}

func calculateMonthlyRepaymentAmount() {
	loanAmount := ask("Enter loan amount", 1000.0)
	interestRate := ask("Enter interest rate", 0.0)
	months := ask("Enter months", 0.2)
	depositAmount := ask("Enter deposit amount", 0.0)

	if depositAmount >= loanAmount {
		fmt.Println("Deposit amount is greater than or same as loan amount.")
		fmt.Println("--")
		calculateMonthlyRepaymentAmount()
	}

	amountToPay := loanAmount - depositAmount	
	monthlyInterestRate := interestRate/100/12
	
	// @see https://www.investopedia.com/terms/a/amortization.asp
	monthlyAmount := amountToPay * ((monthlyInterestRate * math.Pow(1 + monthlyInterestRate, months))/(math.Pow(1 + monthlyInterestRate, months) - 1))

	output := fmt.Sprintf("Monthly repayment amount is %.2f", monthlyAmount)

	displayInfo(output)
}

func ask(label string, min float64) float64 {
	answer := 0.0

	fmt.Printf("%s: ", label)
	fmt.Scanln(&answer)

	if answer < min {
		errorMessage := fmt.Sprintf("value can not be less than %.3f", min)
		fmt.Println(errorMessage)
		ask(label, min)
	}

	return answer
}