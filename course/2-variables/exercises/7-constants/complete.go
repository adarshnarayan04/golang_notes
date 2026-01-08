package main

import "fmt"

func main() {
	// Constants can't use the `:=` short declaration syntax.
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
