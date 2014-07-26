package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's SaturDay?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorro.")
	default:
		fmt.Println("Too far")
	}
}
