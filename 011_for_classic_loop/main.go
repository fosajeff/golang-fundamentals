package main

import "fmt"

func main() {

	// initialization; condition; increment
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	n := 10
	sum := 0

	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}
