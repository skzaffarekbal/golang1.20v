package main

import "fmt"

func main()  {
	fmt.Println("Struct in GoLang")

	// No inheritance and no super or parent in golang
	zaff := User{"Zaff", "zaff@gamil.com", true, 24}
	fmt.Println("Struct OutPut : ", zaff)
	fmt.Printf("Struct OutPut in detail : %+v\n", zaff)
	fmt.Printf("Name : %v\nEmail : %v\n", zaff.Name, zaff.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}