package main


import "fmt"
// The explicit chain — you can trace every error
func Transfer(from, to *Account, amount float64) error {
    if err := from.Withdraw(amount); err != nil {
        return fmt.Errorf("transfer failed: %w", err)
    }
    if err := to.Deposit(amount); err != nil {
        return fmt.Errorf("transfer failed: %w", err)
    }
    return nil
}