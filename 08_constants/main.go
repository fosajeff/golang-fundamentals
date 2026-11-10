package main

import "fmt"

func main() {

	// there are untyped constants and typed constants
	const appName = "Go Basics"

	// typed constants
	const maxUpload int = 100

	const discountedPrice float64 = 19.99

	fmt.Println(appName, maxUpload, discountedPrice)
}
