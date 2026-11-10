package main

import (
	"errors"
	"fmt"
)

func main() {
	// a function called when just before another function returns
	// e.g a cleanup function. defer will guarantee a clean for instance. e.g res.body.Close()
	// e.g again is closing a file

	fmt.Println("Case 1: simulate happy path")
	if err := doWork(true); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Case 2: simulate failure path (fail early)")
	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}
}

func doWork(success bool) error {

	// resource related work
	// start message -> resource acquired
	// cleanup message -> resource released

	fmt.Println("start: resource acquired e.g read a large file")

	// defer will guarantee that this runs at the end of the function no matter the scenario
	// will run if success or error is returned.
	defer fmt.Println("cleanup: resource released")

	if !success {
		return errors.New("Something went wrong. I am returning early")
	}

	fmt.Println("work: doing something important")
	fmt.Println("work: This work is done")

	return nil
}
