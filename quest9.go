package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}
type Product struct{
	Name string
	Price float64
}

func (u *User) HaveBirthday() {
	u.Age++
}

func (u User) Describe() string {
	return fmt.Sprintf("%s (%s) -- %d", u.Name, u.Email, u.Age)
}

type Describer interface {
	Describe() string
}

func printDescription(d Describer) {
	fmt.Print(d.Describe())
}

func printInfo(r Describer){
	fmt.Print(r.Describe())
}
func (p Product) Describe()string{
	return fmt.Sprintf("%s - $%.2f", p.Name, p.Price)
}


func main() {
	user := User{
		Name:  "Michael Bag",
		Email: "michaelbag8@gmail.com",
		Age:   303,
	}
	products := Product{
		Name: "Laptop",
		Price: 999.99,
	}



	fmt.Println("-----Before------")
	fmt.Println(user.Describe())
	user.HaveBirthday()
	fmt.Println("-----After------")
	fmt.Println(user.Describe())
	printDescription(user)

	fmt.Println("---------------")
	fmt.Println()
	printInfo(products)
	printInfo(user)


}