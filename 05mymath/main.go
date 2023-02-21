package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
)

func main() {
	fmt.Println("Learning about Math Package.")

	var numberOne int = 2
	var numberTwo float64 = 4.5

	fmt.Println("The Sum is : ", numberOne+int(numberTwo))

	// Random Number
	// rand.Seed(30)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// Random Number From Crypto
	// rand.Seed(30)
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
