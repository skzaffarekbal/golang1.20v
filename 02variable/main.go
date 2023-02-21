package main

import "fmt"

const LoginToken string = "efrrgtfewe.werrewdef.werffredwf" // always keep const variable name start with capital

func main() {
	var userName string = "Zaff"
	fmt.Println("userName : ", userName)
	fmt.Printf("Variable is type of : %T \n", userName)

	var isLoggedIn bool = false
	fmt.Println("isLoggedIn : ", isLoggedIn)
	fmt.Printf("Variable is type of : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	// var smallVal uint8 = 256 // uint8 ranges from 0-255 it will give overflows error
	fmt.Println("smallVal : ", smallVal)
	fmt.Printf("Variable is type of : %T \n", smallVal)

	var smallFloat float32 = 255.23456787653456
	// var smallFloat float64 = 255.23456787653456
	fmt.Println("smallFloat : ", smallFloat)
	fmt.Printf("Variable is type of : %T \n", smallFloat)

	// default value and some aliases
	var anotherVariable int
	fmt.Println("anotherVariable : ", anotherVariable)
	fmt.Printf("Variable is type of : %T \n", anotherVariable)

	// implicit type
	var website = "google.com"
	fmt.Println("website : ", website)

	// no var style
	numberOfUser := 12 // Can use it inside a method not in global space
	fmt.Println("numberOfUser : ", numberOfUser)

	fmt.Println("LoginToken : ", LoginToken)
	fmt.Printf("Variable is type of : %T \n", LoginToken)
}
