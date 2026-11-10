package main

import "fmt"

// struct groups related fields into one type
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{
		ID:    1,
		Name:  "John Doe",
		Email: "email@gmail.com",
		Age:   10,
	}

	fmt.Println(u1, u1.ID, u1.Email)

	// struct fields are mutable by default
	u1.Age = 20
	fmt.Println(u1)

	u2 := User{
		Name:  "John",
		Email: "john@gmail.com",
	}

	fmt.Println("partial user", u2)
}
