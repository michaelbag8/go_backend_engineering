//Control Flow
package main

var transactions = []float64{500, -200, 0, 100, -1600, 300}
var startingBalance = 1000.00
var owner = "chidi"

func main(){
	for _, eachTrans := range transactions{
		if eachTrans == 0{
			continue
		}
		trans, _:= withdraw(startingBalance,eachTrans)
			if trans == 0{
				break
			}
		
	}
}
