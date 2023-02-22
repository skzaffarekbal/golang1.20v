package main

import "fmt"

func main() {
	var number int = 12
	var result string

	if number > 12 {
		result = "Number is bigger than 12."
	} else if number < 12 {
		result = "Number  is lesser than 12."
	} else {
		result = "Number is equal  to 12."
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
