package main

import(
	"time"
	"fmt"
)

func main() {
	//Q1
	name := "Michael Bag"
	message := "Backend engineering journey begins."
	date := time.Now().Format("2006-01-02")
	fmt.Println(myName(name, message, date))

	//Q2
	Username := "chidi"
	Age := 28
	Balance := 1450.75
	Active := true
	fmt.Println(profile(Username, Age, Balance, Active))


}