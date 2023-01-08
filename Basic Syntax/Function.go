package main

import "fmt"

func suma() {

}
func sumb(a, b int) {

}
func sumc() int {
	return 1
}
func sumd(a, b int) int {
	return a + b
}
func sumAAA(a, b int) (int, string) {
	return a + b, "Hello"
}

func main() {
	suma()
	sumb(1, 2)
	sumc()
	aa := sumd(5, 3)
	aaa, aaa1 := sumAAA(1, 2)
	fmt.Println(aa)
	fmt.Println(aaa, aaa1)
}
