package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sumAndProduct(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}
func main() {
	sum1 := add(1, 3)
	sum2, product := sumAndProduct(3, 4)

	fmt.Println(sum1, sum2, product)

	onlySum, _ := sumAndProduct(4, 5)
	fmt.Println(onlySum)

}
