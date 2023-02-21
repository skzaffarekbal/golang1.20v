package main

import "fmt"

func main()  {
	fmt.Println("Array in GoLang.")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list are : ", fruitList)
	fmt.Println("Fruit list length : ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}

	fmt.Println("Veg list are : ", vegList)
	fmt.Println("Veg list length : ", len(vegList))
}