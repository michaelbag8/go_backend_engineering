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

func filterAge(user []User, age int) []User {
	var output []User
	for _, num := range user {
		if num.Age >= age {
			output = append(output, num)
		}

	}
	return output
}

func main() {
	user := []User{
		{Name: "Mike", Email: "@gmail.com", Age: 30},
		{Name: "James", Email: "@gmail.com", Age: 18},
		{Name: "Coco", Email: "@gmail.com", Age: 34},
		{Name: "Bliss", Email: "@gmail.com", Age: 90},
		{Name: "Bags", Email: "@gmail.com", Age: 60},
		{Name: "Chris", Email: "@gmail.com", Age: 35},
	}
	fmt.Println("-----Original-----")

	for _, each := range user {
		fmt.Println(each.Describe())
	}

	fmt.Println("-----------")

	data := filterAge(user, 20)
	for _, each := range data {
		fmt.Println(each.Describe())
	}
	fmt.Println("-------Filter------")
	for _, each := range user {
		fmt.Println(each.Describe())
	}

}
