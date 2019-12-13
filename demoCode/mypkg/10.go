// !!! THIS PACKAGE SHOULD BE EITHER PUT INTO THE GOPATH
// !!! OR THIS FOLDER SHOULD BE EXPORTED AS GOPATH

package mypkg

import "fmt"

var a notAccessible = "theStuff"

type notAccessible string

// Accessible is just a type wrapper around the built-in string type
type Accessible string

func (n notAccessible) PrintNotAccesible() {
	fmt.Println(n)
}

// PrintAccessible wll print the variable with the "Accessible" type
func (a Accessible) PrintAccessible() {
	fmt.Println(a)
}

// This function cannot be called outside of mypkg
func (a Accessible) printAccessible() {
	fmt.Println(a)
}

// This function cannot be called outside of mypkg
func (n notAccessible) printNotAccesible() {
	fmt.Println(n)
}
