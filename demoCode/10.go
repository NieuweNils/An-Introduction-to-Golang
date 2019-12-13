package main

import "demoIntroGolang/mypkg"

func main() {
	a := mypkg.Accessible("theStuff")
	a.PrintAccessible()
	a.PrintAccessibleANDNotAccessible()

}
