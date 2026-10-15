package main

import "fmt"

func main() {
	// like key value pair in JS
	// map[keyType]valueType
	ages := map[string]int{
		"jeff": 25,
		"john": 35,
	}

	fmt.Println(ages, ages["jeff"], len(ages))

	// make(map[K]V)

	// var scores = map[string]int // nil map
	// fmt.Println(scores, scores["a"]). // error

	scores := make(map[string]int)
	scores["math"] = 90
	fmt.Println(scores)

	users := map[string]string{
		"u1": "jeff",
		"u2": "john",
		"u3": "okoobs",
	}
	fmt.Println(users)

	delete(users, "u2")
	delete(users, "u100") // does not throw any error if you pass invaid key

	fmt.Println(users)
}
