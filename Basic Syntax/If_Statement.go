package main

import "fmt"

func main() {
	// if condition {

	// }
	point := 50
	if point >= 50 {
		fmt.Println("Point >= 50")
	}
	if point >= 40 {
		fmt.Println("Point >= 40")
	} else if point < 60 {
		fmt.Print("50 <= Point < 60 ")
	} else {
		fmt.Println("Else")
	}
}
