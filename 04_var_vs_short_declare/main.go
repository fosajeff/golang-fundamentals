package main

import (
	"fmt"
)

func main() {
	var city string
	city = "London"

	var channel = "Fosa-Tech" // inferred to string

	// := is declare and initialise with type inference

	subscribers := 5000

	subscribers = subscribers + 1000

	likes, comments := 100, 30

	fmt.Println(city, channel, subscribers, likes, comments)
}
