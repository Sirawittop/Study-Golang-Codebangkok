package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}
type Personson struct {
}

//resiver function
func (p Person) Hello() string {
	return "Hello" + p.Name
}

func main() {
	//var x Person
	//xx := Person{}
	xxx := Person{"Sirawit", "Kamchoom", 19}
	xxxx := Person{
		Name:    "Trintrin",
		Surname: "USUSUs",
		Age:     18,
	}
	fmt.Println(xxx.Name, xxx.Surname, xxx.Age)
	fmt.Println(xxxx.Name, xxxx.Surname, xxxx.Age)
	fmt.Println(xxx.Hello())
	fmt.Println(xxxx.Hello())

}
