package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func groupByAge(users []User) map[string][]User {
	var group = make(map[string][]User)
	for _, user := range users {
		if user.Age < 30 {
			group["young"] = append(group["young"], user)
		} else if user.Age >= 30 && user.Age <= 59 {
			group["middle"] = append(group["middle"], user)
		} else {
			group["senior"] = append(group["senior"], user)

		}

	}
	return group
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
	for k, v := range groupByAge(user) {
		for _, each := range v {
			fmt.Println(k, each.Name)
		}

	}

}
