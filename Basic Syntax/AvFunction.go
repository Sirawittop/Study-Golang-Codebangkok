package main

import (
	"fmt"
	"math"
)

//Type Function = (int,int)int
func add(a, b int) int {
	return a + b
}

//Type Function = (int,int)int
func sub(a, b int) int {
	return a - b
}

//Type Function = (string)string
func hello(name string) string {
	return "hello" + name
}

func calculator(f func(int, int) int) {
	sum := f(50, 10)
	fmt.Println(sum)
}
func calculator2(f func(float64, float64) float64) {
	sum := f(50, 2)
	fmt.Println(sum)
}

func sumsumsum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func Varasum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	//Anonymous Function
	x := func(a, b int) int {
		return a + b
	}
	xx := func(a, b int) int {
		return a - b
	}
	xxx := func(a, b int) int {
		return a * b
	}
	xxxx := func(a, b int) int {
		return a / b
	}
	fmt.Println(x(1, 5))
	calculator(x)
	calculator(xx)
	calculator(xxx)
	calculator(xxxx)
	calculator(add)
	calculator(sub)
	calculator2(func(i1, i2 float64) float64 { return (math.Pow(i1, i2)) })
	value := []int{1, 2, 3, 4, 5, 6}
	sumofvalue := sumsumsum(value)
	fmt.Println("Sum of Value : ", sumofvalue)
	////
	sumvadafunc := Varasum(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum of Value : ", sumvadafunc)
}
