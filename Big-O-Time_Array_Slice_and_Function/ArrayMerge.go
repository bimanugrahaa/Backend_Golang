package main

import "fmt"

//Merge the same array values
func ArrayMerge(arrayA, arrayB []string) []string {

	array := append(arrayA, arrayB...)
	check := make(map[string]bool)
	list := []string{}

	for _, val := range array {
		value := check[val]
		if !value {
			check[val] = true
			list = append(list, val)
		}
	}

	return list
}

func main() {

	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []

}
