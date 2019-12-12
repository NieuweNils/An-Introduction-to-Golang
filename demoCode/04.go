package main

import (
	"fmt"
)

func main() {
	numbers := [6]int{3, 16, -2, 10, 23, 12}
	intSlice := []int{}
	for i, number := range numbers {
		if number >= 10 && number <= 20 {
			intSlice = append(intSlice, i)
		}
	}

	fmt.Println("Indeces with a value of between 10 and 20:")
	for _, arrayIndex := range intSlice {
		fmt.Printf("index %d has number: %d\n", arrayIndex, numbers[arrayIndex])
	}
}
