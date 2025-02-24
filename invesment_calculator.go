package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	var invesmentAmount float64
	expectedReturnRate := 5.5
	var years float64 = 10

	fmt.Print("Invesment Amount: ")
	fmt.Scan(&invesmentAmount)
	fmt.Print("Expected Return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("No of Years: ")
	fmt.Scan(&years)

	futureValue := invesmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futRealValue)

}
