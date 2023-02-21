package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Learning about Time Package.")

	presentTime := time.Now()
	fmt.Println("presentTime : ", presentTime)

	fmt.Println("presentTime time format : ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.May, 11, 23, 23, 15, 0, time.UTC)
	fmt.Println("createdDate : ", createdDate)
	fmt.Println("createdDate time format : ", createdDate.Format("01-02-2006 15:04:05 Monday"))
}