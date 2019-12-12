package main

import (
	"fmt"
	"reflect"
)

// Population is just a type wrapper around the built-in Integer type
type Population int

func main() {
	var population Population = 572
	fmt.Println("type of \"var population Population= 572\":", reflect.TypeOf(population))

	population2 := Population(572)
	fmt.Println("type of \"population2 := Population(572)\":", reflect.TypeOf(population2), "\n")

	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println("Congrats, Kevin and Anna! It's a girl!")
	population++
	fmt.Println("Sleepy Creek County population:", population)
}
