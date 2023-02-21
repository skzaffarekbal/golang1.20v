package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating for our pizza : ")

	// comma ok || err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating : ", input)

	fmt.Printf("Type of input : %T\n", input)
}
