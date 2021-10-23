package main

import "fmt"

type Student struct {
	name []string

	score []int
}

func (s Student) Average() float64 {

	// fmt.Println(mapStudent)
	avg := 0.0
	num := 0
	for key, val := range s.score {
		num += val
		avg = float64(num / int(key+1))
	}

	return avg
}

func (s Student) Min() (min int, name string) {

	min = s.score[0]

	for key, num := range s.score {

		if num < min {
			name = s.name[key]
			min = num
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {

	max = s.score[0]

	for key, num := range s.score {

		if num > max {
			name = s.name[key]
			max = num
		}
	}

	return max, name
}

func main() {

	var a = Student{}

	// mapStudent := make(map[string]int)
	for i := 0; i < 5; i++ {

		var name string

		fmt.Print("Input " + string(i) + " Student’s Name : ")

		fmt.Scan(&name)

		a.name = append(a.name, name)

		var score int

		fmt.Print("Input " + name + " Score : ")

		fmt.Scan(&score)

		a.score = append(a.score, score)

		// mapStudent[name] = score
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())

	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

}
