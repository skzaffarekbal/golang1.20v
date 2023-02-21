package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Slices in GoLang.")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana");
	fmt.Println("Fruit List : ", fruitList)

	// fruitList = append(fruitList[1:]);
	// fruitList = append(fruitList[:3]);
	fruitList = append(fruitList[1:3]);
	fmt.Println("Fruit List : ", fruitList)

	highScores := make([]int, 4 )

	highScores[0] = 234
	highScores[1] = 245
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 999 // it will gives error due to range

	highScores = append(highScores, 555, 666, 777)

	fmt.Println("HighScores : ", highScores);

	fmt.Println("is all data sorted : ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("HighScores : ", highScores);
	fmt.Println("is all data sorted : ", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var courses = []string{"react.js", "javascript", "java", "swift", "go"}
	fmt.Println("Courses : ", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses : ", courses)
}