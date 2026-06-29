package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) Describe() string {
	return fmt.Sprintf("%s (%s) -- %d", u.Name, u.Email, u.Age)
}

func main() {
	user := User{
		Name:  "Michael Bag",
		Email: "michaelbag8@gmail.com",
		Age:   303,
	}
	//method call
	fmt.Println(user.Describe())
}
