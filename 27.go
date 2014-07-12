package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1,2}
	r := p
	q := &p
	q.X = 1e9


	fmt.Println(p)
	fmt.Println(r)
}
