package main

import "fmt"

//Get the fibonacci number using bottom up
func fibo(n int) int {

	// your code here

	//Implementation with add the last two fibo number
	// a := 0   //Fibo 0
	// b := 1   //Fibo 1
	// i := 2 //If n > 1, start at n =2
	// now := 0 //Fibo now (total)

	// if n == 0 || n == 1 {
	// 	return n
	// }

	// for i <= n {
	// 	now = a + b
	// 	a, b = b, now
	// 	i++
	// }

	//Implementation using map
	i := 2                       //If n > 1, start at n =2
	mapFibo := make(map[int]int) //Map the fibo number

	mapFibo[0] = 0 //Fibo 0
	mapFibo[1] = 1 //Fibo 1

	for i <= n {
		mapFibo[i] = mapFibo[i-1] + mapFibo[i-2]
		i++
	}

	return mapFibo[n]

}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 13

	fmt.Println(fibo(10)) // 55

}
