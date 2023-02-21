package main

import "fmt"

func main()  {
	fmt.Println("Map in GoLang")
	
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "GoLang"

	fmt.Println("languages map data : ", languages)
	fmt.Println("JS short for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("languages map data : ", languages)

	// Loops are interesting in golang
	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}