package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5
	//%v - general value ( can be used for any type)

	//.1f -> round to 1 decimal place
	//.2f -> round to 2 decimal places

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	// don't edit below this line

	fmt.Println(msg)
}
