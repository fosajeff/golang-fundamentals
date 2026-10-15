package main

import "fmt"

func main() {
	// arrays are fixed sizes in go and cannot grow
	var marks [3]int
	marks[0] = 10
	marks[1] = 20
	marks[2] = 50

	fmt.Println(marks)

	// array literal -> quick way to create a fixed size array
	res := [5]int{2, 3, 4, 5, 6}

	fmt.Println(len(res))
}
