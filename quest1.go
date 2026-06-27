// Understanding Go toolchain
package main

import (
	"fmt"
	"time"
)

func myName(name, message string, date string) (string, string, time.Time) {
	return name, message, time.Now()
}
func main() {
	name := "Michael Bag"
	message := "Backend engineering journey begins."
	date := time.Now().Format("2006-01-02")
	fmt.Println(myName(name, message, date))

}
