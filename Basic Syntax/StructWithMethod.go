package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) GetName() string {
	return p.name
}
func (p Person) GetAge() int {
	return p.age
}
func (p *Person) SetName(name string) {
	p.name = name
}
func (p *Person) setAge(age int) {
	p.age = age
}

func main() {
	x := Person{"Sirawit", 19}
	fmt.Println(x.name, x.age)
	x.SetName("TrinTrin")
	x.setAge(20)
	fmt.Println(x.name, x.age)
}
