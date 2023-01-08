package main

import "fmt"

func sum(result *int) {
	a := 10
	b := 50
	*result = a + b
}
func main() {
	x := 10
	y := x
	y = 20 // x  != 20
	fmt.Println(&x)
	fmt.Println(y)

	var a, b int
	a = 10
	b = a
	fmt.Println(&a)
	fmt.Println(&b)

	//same address
	var i int
	i = 20
	var t *int
	t = &i
	fmt.Println(&i)
	fmt.Println(t)
	fmt.Println(*t)
	fmt.Println(i)
	*t = 30
	fmt.Println("____")
	fmt.Println(*t)
	fmt.Println(i)
	var result int
	sum(&result)
	fmt.Println("resut:", result)
	// aa := sum(&hh)
	// fmt.Println(aa)
}
