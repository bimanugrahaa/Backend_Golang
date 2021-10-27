package main

import "fmt"

//Get coins from given money
func moneyCoins(money int) []int {

	// your code here

	arrMoney := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	i := len(arrMoney) - 1
	arrCoinsM := []int{}

	for money > 0 {
		if money >= arrMoney[i] {
			money -= arrMoney[i]
			arrCoinsM = append(arrCoinsM, arrMoney[i])
		} else {
			i--
		}

	}
	return arrCoinsM
}

func main() {

	fmt.Println(moneyCoins(123)) // [100 20 1 1 1]

	fmt.Println(moneyCoins(432)) // [200 200 20 10 1 1]

	fmt.Println(moneyCoins(543)) // [500, 20, 20, 1, 1, 1]

	fmt.Println(moneyCoins(7752)) // [5000, 2000, 500, 200, 50, 1, 1]

	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]

}
