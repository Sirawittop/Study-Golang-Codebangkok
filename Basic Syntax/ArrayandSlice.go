package main

import "fmt"

func ArrayUnlimet() {
	x := [...]int{1, 2, 3, 4, 5, 999}
	fmt.Println(x)
}

func SliceList() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	x = append(x, 5)
	fmt.Println(x)
	y := len(x)
	fmt.Println(y)
	xx := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	xxx := xx[4:7] // end index+1
	fmt.Println(xxx)
}

func ArrayLimit() {
	x := [5]int{1, 2} // Zero Value = 0
	fmt.Println(x)
}

func Array2D() {
	x := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(x)
	x[0][0] = 500
	fmt.Println(x)
}

func main() {
	var x [3]int = [3]int{1, 2, 3}
	var y [5]int = [5]int{}
	xx := [3]int{2, 3, 4}
	yy := [5]int{}
	fmt.Println(x, y, xx, yy)
	Array2D()
	ArrayLimit()
	ArrayUnlimet()
	SliceList()
}
