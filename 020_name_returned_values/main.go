package main

import "fmt"

func divide(a int, b int) (result int, reminder int) {
	result = a / b
	reminder = a % b
	return
}

func main() {
	res, rem := divide(20, 3)
	fmt.Println(res, rem)
}
