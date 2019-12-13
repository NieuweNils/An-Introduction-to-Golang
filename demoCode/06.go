package main

import (
	"fmt"
	"reflect"
)

// Population is just a type wrapper around the built-in Integer type
type Population int

// Date is a struct type that holds a year, month, and day value.
type Date struct {
	year  int
	month int
	day   int
}

// PopulationHistory is a type that embeds the Date and Population types
type PopulationHistory struct {
	date       Date
	population Population
}

func main() {
	populationHistory := []PopulationHistory{}
	today := Date{2019, 12, 13}
	populationToday := PopulationHistory{today, 572}
	populationHistory = append(populationHistory, populationToday)

	fmt.Println("type of \"population\":", reflect.TypeOf(populationHistory), "\n")
	fmt.Println("date today:", populationToday.date)

	fmt.Println("Sleepy Creek County population history:", populationHistory, "\n")
	fmt.Println("Your day breaks.. your mind aches.. You find that all.. ")
	populationToday.date.day++
	fmt.Println("date:", populationToday.date)
	fmt.Println("population:", populationToday.population, "\n")

	fmt.Println("Congrats, Kevin and Anna! It's a girl!")
	populationToday.population++
	fmt.Println("population:", populationToday.population, "\n")
	populationHistory = append(populationHistory, populationToday)

	fmt.Println("Sleepy Creek County population history:")
	for _, history := range populationHistory {
		fmt.Printf("date: %d-%d-%d \t population: %d\n", history.date.year, history.date.month, history.date.day, history.population)
	}

	fmt.Println(populationHistory)
	for _, history := range populationHistory {
		fmt.Println(history)
	}
}
