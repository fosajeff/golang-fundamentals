package main

import "fmt"

func main() {
	// stores the mem address of a variable
	// two things to know: operators and address
	// &x -> address of x (makes a pointer)
	// *p -> dereference (go to that address and read/write)
	// if you want to change values in a func without returning it, this applies

	score := 10
	fmt.Println("before:", score)

	addScore(&score)
	fmt.Println("after:", score)

}

func addScore(s *int) {
	// read and write
	*s = *s + 5
}
