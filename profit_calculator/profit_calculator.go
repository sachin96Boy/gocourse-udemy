package main

import "fmt"

func main() {

	revenue := 0.0
	expence := 0.0
	tax_rate := 0.0

	handleInput(&revenue, &expence, &tax_rate)

	ebt, profit, ratio := handleCalcuulation(revenue, expence, tax_rate)

	printResult(ebt, profit, ratio)

}

func handleInput(revenue, expence, tax_rate *float64) {
	fmt.Print("Total Revenue: ")
	fmt.Scan(revenue)
	fmt.Print("Total Expence: ")
	fmt.Scan(expence)
	fmt.Print("Tax rate: ")
	fmt.Scan(tax_rate)
}

func handleCalcuulation(revenue, expence, tax_rate float64) (float64, float64, float64) {

	ebt := revenue - expence
	profit := ebt * (1 - tax_rate/100)

	ratio := ebt / profit

	return ebt, profit, ratio

}

func printResult(ebt, profit, ratio float64) {
	fmt.Printf("Earnings Before Tax: %.2f\n", ebt)
	fmt.Printf("Profit (Earning After Tax): %.2f\n", profit)
	fmt.Printf("ratio %.3f\n", ratio)
}
