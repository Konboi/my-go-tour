package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The vale:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The vale:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The vale:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
