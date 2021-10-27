package main

import (
	"fmt"
	"sort"
)

type pair struct {
	key   string
	value int
}

//Count items from given string
func MostAppearItem(items []string) []pair {

	// your code here

	mapItem := make(map[string]int)
	p := make([]pair, 0)

	//Add the value if there is duplicate item
	for _, val := range items {
		mapItem[val]++
	}

	//Append map to pair
	for key, value := range mapItem {
		p = append(p, pair{key, value})
	}

	//Sort pair ascending
	sort.Slice(p, func(i, j int) bool {
		// fmt.Printf("%d\t%d\n", p[i].value, p[j].value)
		// fmt.Println(p[i].value < p[j].value)
		return p[i].value < p[j].value
	})

	return p

}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))

	// golang->1 ruby->2 js->4

	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))

	// C->1 D->2 B->3 A->4

	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))

	// football->1 basketball->1 tenis->1

}
