package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
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
		{Name: "Mike", Age: 30},
		{Name: "James", Age: 18},
		{Name: "Coco", Age: 34},
		{Name: "Bliss", Age: 90},
		{Name: "Bag", Age: 60},
		{Name: "Chris", Age: 35},
	}
	fmt.Println("-----Original-----")

	fmt.Println(user)

	fmt.Println()

	fmt.Println(filterAge(user, 20))
	fmt.Println("-------Filter------")
	fmt.Println(user)

}
