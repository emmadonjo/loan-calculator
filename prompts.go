package main

import (
	"fmt"
)

func Run() {
	selectAction := action()

    switch selectAction {
		case "1":
			CalculateMonthlyRepaymentAmount()
		case "5":
			return
		default:
			Run()
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

func Ask(label string, min float64) float64 {
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


func DisplayInfo(info string) {
	fmt.Println()
	fmt.Println(info)
	fmt.Println("-----------------------------------")
	fmt.Println()
	Run()
}