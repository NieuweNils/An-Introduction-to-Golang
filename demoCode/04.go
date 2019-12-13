package main

import (
	"fmt"
)

func main() {
	numbers := [6]int{3, 16, -2, 10, 23, 12}
	intSlice := []int{}
	for _, number := range numbers {
		if number >= 10 &&
			number <= 20 {
			intSlice = append(intSlice, number)
		}
	}

	fmt.Println("This slice contains the following values of between 10 and 20:")
	for _, value := range intSlice {
		fmt.Println(value)
	}
}
