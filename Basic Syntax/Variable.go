package main

import "fmt"

func ShortDeclaced() {
	variableone := 220
	_ = variableone
}

func ZeroValue() {
	var x int
	var y string
	var z float64
	var xx bool
	fmt.Println(x, y, z, xx)

}
func main() {
	var variableone int
	const variabletwo = "sirawit"
	variableone = 220
	_ = variableone
	ZeroValue()
}

//Short Declaced = Local Variable (Infunction)
//Var Declaced   = Gobal Variable (Outfunction)
// _ = variable (Ignore)
