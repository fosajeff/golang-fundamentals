package main

import "fmt"

func main() {
	// make([]type, length, capacity)
	// this is used for performance
	scores := make([]int, 0, 5)

	fmt.Println(scores, len(scores), cap(scores))

	scores = append(scores, 100)
	fmt.Println("after appending 100", scores)
	scores = append(scores, 200, 3000)
	fmt.Println("After appenindg 200, 3000", scores)

	scores = append(scores, 45, 55)
	fmt.Println("After appending 45, 55", scores)

	// if in case we are exceeding the capacity, go grows the backing array (usually doubles)
	scores = append(scores, 60)
	fmt.Println("After appending 60", scores, len(scores), cap(scores))

	todos := []string{"Do youtube", "workout everyday"}

	more := []string{"learn golang"}

	// ... spread elements of a slice into another
	todos = append(todos, more...)
	fmt.Println(todos)
}
