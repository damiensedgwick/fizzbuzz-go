package main

import (
	"fmt"
)

func main() {
	var stopAt uint32 = 1000000

	for i := uint32(1); i <= stopAt; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}

		if i%5 == 0 {
			fmt.Print("Buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i)
		}

		fmt.Println()
	}
}
