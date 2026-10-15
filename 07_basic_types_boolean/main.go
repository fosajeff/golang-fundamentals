package main

import (
	"fmt"
)

func main() {
	isLogged := true // inferred as boolean type
	isAdmin := true
	hasSubscription := false

	// AND &&
	canOpenDashboard := isLogged && hasSubscription

	// OR ||
	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println(canOpenDashboard, canDeletePost)

	age := 25
	isAdult := age > 18
	fmt.Println(isAdult)
}
