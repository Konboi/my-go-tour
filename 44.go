package main

import "fmt"

func fibonacci() func () int {
	fibo := 1
	before_fibo := 0
	return func() int {
		bbb := fibo
		fibo = before_fibo + fibo
		before_fibo = bbb
		return fibo
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
