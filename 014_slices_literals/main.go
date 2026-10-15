package main

import "fmt"

func main() {
	// slices are most common collection type.
	// They can be dynamic and can grow
	// []type{...}

	results := []string{"Jeffrey", "Efosa"}
	fmt.Println(results, results[len(results)-1])

	results[1] = "Okooboh"
	fmt.Println(results)

	var nums []int

	nums = append(nums, 10)
	nums = append(nums, 20, 30)

	fmt.Println(nums)
}
