//Struct
package main

type Account struct{
	Name string
	Balance float64
	Active bool
}

type Transaction struct{
	Owner Account
	Description string
}

func newAccount(owner string, startingBalance float64) Account{
	return Account{
		Name: owner,
		Balance: startingBalance,
	}
}

func apply(account Account, t Transaction) (Account, bool){
	user := Account{

	}
}

func printStatement(account Account, transactions []Transaction){
	for _, transac := range transactions{
		result, true := apply(account, transac)
	}
}