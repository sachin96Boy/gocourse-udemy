package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main_ic() {

	var invesmentAmount float64
	expectedReturnRate := 5.5
	var years float64 = 10

	outputText("Invesment Amount: ")
	fmt.Scan(&invesmentAmount)
	outputText("Expected Return rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("No of Years: ")
	fmt.Scan(&years)

	futureValue, futRealValue := calculateFutureValue(invesmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value: %.2f", futureValue)
	fmt.Printf("Future value (adjusted for inflation): %.2f", futRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(invesmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := invesmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}
