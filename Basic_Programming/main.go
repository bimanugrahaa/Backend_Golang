package main

import (
	"fmt"
)

func main() {
	// prob1()
	// prob2()
	// prob3()
	prob4()
	// prob5()
	// prob6()
	// prob7()
	// prob8()
}

func prob1() {
	fmt.Println("=========== 1 LP Tabung ===========")
	const pi float64 = 3.14
	var r, T, lp float64

	fmt.Print("Masukan jari-jari tabung = ")
	fmt.Scanf("%2f\n", &r)
	fmt.Print("Masukan tinggi tabung = ")
	fmt.Scanf("%2f\n", &T)

	lp = (2 * pi * r) * (r + T)
	fmt.Print("Luas permukaan tabung = ")
	fmt.Println(lp)
}

func prob2() {
	fmt.Println("=========== 2 Grade Nilai ===========")

	var name, grade string
	var score int

	fmt.Print("Masukan nama = ")
	fmt.Scanf("%s \n", &name)

	fmt.Print("Masukan nilai = ")
	fmt.Scanln(&score)

	switch {
	case score <= 100 && score >= 80:
		grade = "A"
	case score <= 79 && score >= 65:
		grade = "B"
	case score <= 64 && score >= 50:
		grade = "C"
	case score <= 49 && score >= 35:
		grade = "D"
	case score <= 34 && score >= 0:
		grade = "E"
	default:
		grade = "Invalid"
	}

	fmt.Printf("Nilai %s = %s \n", name, grade)
}

func prob3() {
	fmt.Println("=========== 3 Faktor Bilangan ===========")

	var number int

	fmt.Print("Masukan bilangan = ")
	fmt.Scanf("%v \n", &number)

	fmt.Printf("Faktor dari %v adalah ", number)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Printf("%v ", i)
		}
	}
}

func prob4() {
	fmt.Println("=========== 4 Bilangan Prima ===========")

	var number int

	fmt.Print("Masukkan bilangan = ")
	fmt.Scanf("%v \n", &number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			if i != 1 && i != number {
				fmt.Println("Bukan Bilangan Prima")
				break
			}

			if i == number {
				fmt.Println("Bilangan Prima")
			}
		}
	}
}

func prob5() {
	fmt.Println("=========== 5 Palindrome ===========")

	var str string
	var newStr string
	var word int
	var sentences string

	fmt.Print("Masukkan jumlah kata = ")
	fmt.Scan(&word)

	fmt.Printf("Masukkan %v kata = ", word)
	for i := 0; i < word; i++ {
		fmt.Scan(&str)
		sentences += str
	}

	for i := len(sentences) - 1; i >= 0; i-- {
		newStr += string(sentences[i])
	}

	if sentences == newStr {
		fmt.Print("That word is Palindrome \n")
	} else {
		fmt.Print("That word isn't Palindrome \n")
	}
}

func prob6() {
	fmt.Println("=========== 6 Exponentiation ===========")

	var x int
	var n int
	result := 1

	fmt.Print("Masukkan angka = ")
	fmt.Scan(&x)
	fmt.Print("Masukkan pangkat = ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		result = result * x
	}

	fmt.Println(result)
}

func prob7() {
	fmt.Println("=========== 7 Asterisk ===========")

	var rows int

	fmt.Print("Masukkan jumlah baris = ")
	fmt.Scanf("%v \n", &rows)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= (2 * i); k++ {
			if k%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Println()
	}
}

func prob8() {
	fmt.Println("=========== 8 Perkalian ===========")

	var number int

	fmt.Print("Masukkan bilangan = ")
	fmt.Scanf("%v \n", &number)

	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			result := i * j
			fmt.Printf("%v \t", result)
		}
		fmt.Println()
	}
}
