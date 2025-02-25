package main

import "fmt"

func main() {

	revenue := 0.0
	expence := 0.0
	tax_rate := 0.0

	fmt.Print("Total Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Total Expence: ")
	fmt.Scan(&expence)
	fmt.Print("Tax rate: ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expence
	profit := ebt * (1 - tax_rate/100)

	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
