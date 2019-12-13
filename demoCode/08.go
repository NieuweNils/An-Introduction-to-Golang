package main

import "fmt"

// Population is just a type wrapper around the built-in Integer type
type Population int

// Decimal is just a type wrapper around the built-in Float64 type
type Decimal float64

// Add is a method that will add the specified number of people to the Population count it is called on.
func (p *Population) Add(newPersonAmount int) {
	*p += Population(newPersonAmount)
}

// Add is a method that will add the specified number to the Decimal it is called on.
func (d *Decimal) Add(number int) {
	*d += Decimal(number)
}

// Addable is an interface that will be available to any type that has an "Add" method
type Addable interface {
	Add(number int)
}

func main() {
	population := Population(0)
	decimal := Decimal(0)

	// Population & Decimal satisfy the Addable interface and
	// have methods with pointer receivers
	var popPoint Addable = &population
	var decPoint Addable = &decimal

	fmt.Println("population before:", population)
	fmt.Println("decimal before:", decimal)

	popPoint.Add(3)
	decPoint.Add(3)
	popPoint.Add2(2)

	fmt.Println("population after:", population)
	fmt.Println("decimal after:", decimal)
}

// Add2 is a method that will add the specified number of people to the Population count it is called on.
func (p *Population) Add2(newPersonAmount int) {
	*p += Population(newPersonAmount)
}
