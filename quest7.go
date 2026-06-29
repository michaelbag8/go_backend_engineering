package main

import (
	"fmt"
)

type Users struct {
	Name  string
	Email string
	Age   int
}

func (u *Users) HaveBirthday(){
	u.Age++
}

func main() {
	user := Users{
		Name:  "Michael Bag",
		Email: "michaelbag8@gmail.com",
		Age:   303,
	}
	//method call
	user.HaveBirthday()
}
