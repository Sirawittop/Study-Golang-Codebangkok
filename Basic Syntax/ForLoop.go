package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	//for loop
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	//while loop
	a := 0
	for a < len(x) {
		fmt.Println(x[a])
		a++
	}
	//for length
	for i, v := range x {
		fmt.Println(i, v)
	}
	//for length enor
	for _, v := range x {
		fmt.Println(v)
	}
}
