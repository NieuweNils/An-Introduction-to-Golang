package main

import "fmt"

// Population is just a type wrapper around the built-in Integer type
type Population int

// Add is a method that will add the specified number of people to the Population count it is called on.
func (p *Population) Add(newPersonAmount int) {
	if newPersonAmount < 0 {
		fmt.Println("negative babies can't be born, stupid.")
	} else {
		*p += Population(newPersonAmount)
	}
}

func main() {
	population := Population(2)
	fmt.Println("Population before:", population, "\n")
	babies := [6]int{2, 5, 6, -1, 87, 2}
	for _, babyAmount := range babies {
		fmt.Println("BAM!", babyAmount, "babies were born just now!")
		population.Add(babyAmount)
		fmt.Println("the population is now equal to:", population, "\n")
	}
}
