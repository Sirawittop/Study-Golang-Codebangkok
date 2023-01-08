package main

import "fmt"

func main() {
	var coutries map[string]string
	_ = coutries
	courestwo := map[string]string{}
	courestwo["TH"] = "Thailand"
	courestwo["EN"] = "England"
	fmt.Println(courestwo["TH"])
	coutriestwo, ok := courestwo["eJS"]
	if ok {
		fmt.Println(coutriestwo)
	} else {
		fmt.Println("No value")
	}
}
