package main

import "fmt"

func main() {
	fmt.Println("Pointers in GoLang")

	var ptr *int
	fmt.Println("ptr value : ", ptr)

	myNumber := 23
	var ptr1 = &myNumber

	fmt.Println("value of data and pointer address : ", *ptr1, ptr1 )

	*ptr1 = *ptr1 * 2
	fmt.Println("Value : ", myNumber)
}
