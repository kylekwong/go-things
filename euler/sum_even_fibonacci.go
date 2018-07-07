package main

import "fmt"

// Find the sum of all even Fibonacci numbers < 4,000,000
func main() {
	sum, i, j := 0, 0, 1

	for j < 4000000 {
		if j%2 == 0 {
			sum += j
		}
		i, j = j, i+j
	}
	fmt.Println("Sum: ", sum)
}
