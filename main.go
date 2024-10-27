package main

import (
	"fmt"
	"time"
)

func fizzBuzz(stopAt uint32) {
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

func runBenchmark(stopAt uint32) float64 {
	start := time.Now()
	fizzBuzz(stopAt)
	duration := time.Since(start)
	return duration.Seconds()
}

func main() {
	const stopAt uint32 = 1000000
	seconds := runBenchmark(stopAt)
	fmt.Printf("\nTime taken: %.6fs\n", seconds)
}

func BenchmarkFizzBuzz(stopAt uint32) float64 {
	return runBenchmark(stopAt)
}
