// Functions and Multiple Return Values
package main

import "fmt"

func deposit(balance float64, amount float64) float64{
	new_balance := balance + amount
	return new_balance
}

func withdraw(balance float64, amount float64) (float64, bool){
	if amount > balance{
		return 0, false
	}
	return balance-amount, true
}

func summarize(owner string, balance float64) string{
	return fmt.Sprintf("%s, your bank balance: %.2f",owner, balance)
}

