package main

import "fmt"

func greatest(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest

}

func main() {
	z := greatest(14, 25, 98, 77, 85, 23, 46)

	fmt.Println(z)
}
