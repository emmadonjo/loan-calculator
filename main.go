package main

import (
	"fmt"
)

func main() {
	Run()	
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