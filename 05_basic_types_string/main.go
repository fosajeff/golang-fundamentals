package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Efosa"
	lastName := "Okooboh"
	fullName := firstName + " " + lastName

	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}
