package main

import (
	"fmt"
	"math"
)
func Sqrt(x float64)  float64 {
	var before_z float64
	var z float64 = 1.0

	for	i := 1; i < 10000; i++ {
		before_z = z
		z = z - ((z*z) - x) / 2*z

		if 0.000001 > math.Abs(z - before_z) {
			i = 1000000000
		}
	}

	return z
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
